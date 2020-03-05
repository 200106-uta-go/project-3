package kreate

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

/*
## kreate run <profile name>
1. Prerequisite: Helm v2.16.3 must be installed and helm init must be ran (confirm)
2. Check: If the istio environment is already deployed to the cluster, there is no need to redeploy it.
    - If istio is not deployed, istio v1.4.5 must be deployed.
3. The custom chart must be created using Part 1. of the the 'Kreate chart' logic.
4. Check: If a custom chart is already deployed to the cluster, it must be upgraded to the new chart.
    - If a custom chart is not already deployed, the new chart must be installed rather than upgraded.
	- Possible lead for the implementing developer: https://github.com/helm/helm/issues/3353
*/

func main() {
	fmt.Println(RunProfile("./test.yaml"))
}

func RunProfile(profileName string) string { // current logic was written prior to the 3/3/20 MVP meeting
	currentDir := "./"

	/*
		// 1. Prerequisite: Confirm Helm v2.16.4 is installed. JZ: I think we should hang on to this logic since initializing helm v2.16.2/deploying istio might be relocated to init in the future.
		str, err := shellCommand("sudo helm version", currentDir)
		if !strings.Contains(str, "v2.16.3") && err == nil {
			return fmt.Sprintf("Error: Helm version is not v2.16.3 (required). Please confirm your Helm version.")
		} else if err != nil {
			// Prerequisite: Confirm Helm init is already ran.
			if strings.Contains(str, "could not find tiller") {
				return fmt.Sprint("Error: Could not find tiller. Please confirm that Helm is initialized.")
			}
			// Misc. error (Helm not installed, no Cluster, ect.)
			return fmt.Sprintf("Error: %s", str)
		}
	*/

	// 1. Confirm Helm conditions
	str, err := shellCommand("sudo helm version", currentDir)

	// Check if v2.16.3 is installed
	if !strings.Contains(str, "v2.16.3") && err == nil {
		fmt.Println("Error: Helm version is not v2.16.3 (required). Attempting to run install script...")
		if runInstallScript() != nil {
			return fmt.Sprint("Error: Could not resolve Helm version mismatch.")
		}

	} else if err != nil {

		// Check if tiller is not initialized yet
		if strings.Contains(str, "could not find tiller") {
			fmt.Println("Error: Could not find tiller. Attempting to run install script...")
			shellCommand("kubectl -n kube-system delete deployment tiller-deploy", currentDir)
			shellCommand("kubectl -n kube-system delete service/tiller-deploy", currentDir)
			if runInstallScript() != nil {
				return fmt.Sprint("Error: Could not resolve missing tiller.")
			}

			// Misc. error occurred (Helm not installed, no Cluster, ect.)
			fmt.Printf("Error: %s. Attempting to run install script...\n", str)
			if runInstallScript() != nil {
				return fmt.Sprintf("Error: Could not resolve: %s.", str)
			}
		}
	}

	// 2. Check if istio is already deployed
	str, err = shellCommand("kubectl get services -n istio-system", currentDir)

	// Check for deployed services
	if strings.Contains(str, "No resources found in istio-system namespace") {
		fmt.Println("Istio is not yet deployed. Attempting to run install script...")
		if runInstallScript() != nil {
			return fmt.Sprint("Error: Failed to deploy Istio.")
		}

	} else if err != nil {

		// Misc. error occurred (Anomaly, hoping a scripted reinstall will fix it?)
		fmt.Printf("An error occurred when querying the Kubernetes cluster: %s Attempting to rerun install script...\n", str)
		if runInstallScript() != nil {
			return fmt.Sprintf("Error: Could not resolve %s.", str)
		}
	}
	//================================================================
	// 3. Create the custom chart
	profile := GetProfile(PROFILES + profileName + ".yaml")

	// Creates a new values.yml & Chart.yml file.
	createValues(profile)
	createChartFile(profile)

	//build file structure for running helm
	buildFileSystem()

	//add values into chart for deployment yaml
	populateChart("values.yaml", "./templates")
	//==============================================================

	// 4. Deploy/Upgrade custom chart
	return fmt.Sprintf("Profile %s deployed successfully", profileName)
}

func shellCommand(command, dir string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func runInstallScript() error {
	const filename = "tempSetup.sh"

	const installScript = `sudo curl -L https://github.com/istio/istio/releases/download/1.4.5/istio-1.4.5-linux.tar.gz -o istio-1.4.5.tar.gz
	tar -xf istio-1.4.5.tar.gz
	cd istio-1.4.5
	export PATH=$PWD/bin:$PATH
	sudo curl -L https://get.helm.sh/helm-v2.16.3-linux-amd64.tar.gz -o helm.tar.gz
	tar -xf helm.tar.gz
	cd linux-amd64/
	sudo cp helm /bin/helm
	sudo cp tiller /bin/tiller
	cd ..
	sudo kubectl apply -f install/kubernetes/helm/helm-service-account.yaml
	sudo helm init --service-account tiller
	echo "waiting for tiller pod to be ready ..."
	sudo kubectl -n kube-system wait --for=condition=Ready pod -l name=tiller --timeout=300s
	sudo helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
	echo "waiting for istio-system jobs to complete (will take about minute)"
	kubectl -n istio-system wait --for=condition=complete job --all
	sudo helm install install/kubernetes/helm/istio --name istio --namespace istio-system --values install/kubernetes/helm/istio/values-istio-demo.yaml` // script goes here

	er := ioutil.WriteFile(filename, []byte(installScript), 0777)
	if er != nil {
		return er
	}
	cmd := exec.Command("/bin/sh", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	er = cmd.Run()
	os.Remove(filename)
	return er
}
