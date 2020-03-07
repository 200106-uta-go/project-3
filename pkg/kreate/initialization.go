package kreate

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// This value will determine where the helm directories will go by default.
const (
	MOULDFOLDERS = "/var/local/kreate/" // Initial value on where to store helm charts
	PROFILES     = "/etc/kreate/"
)

var (
	chartsLocation string
)

/*
## kreate init
1. Setup kreate's folders to the proper paths (var/local/kreate holds the istio and custom moulds. etc/kreate/ holds profile .yaml files.)
2. Setup kreate's environment variables (If any)
*/

func InitializeDirectories() {
	pathErr := os.MkdirAll(MOULDFOLDERS, 0777)
	if pathErr != nil {
		log.Panicf("Error making directory %s => %v", MOULDFOLDERS, pathErr)
	}

	pathErr = os.MkdirAll(PROFILES, 0777)
	if pathErr != nil {
		log.Panicf("Error making directory %s => %v", PROFILES, pathErr)
	}

	// Register the path of where the helm charts will be stored by checking the env var "KREATE_DATA"
	// Set "KREATE_DATA" == The default value of /var/local/kreate/
	var ok bool
	chartsLocation, ok = os.LookupEnv("KREATE_DATA")
	if !ok {
		setErr := os.Setenv("KREATE_DATA", MOULDFOLDERS)
		if setErr != nil {
			log.Panicf("Error Setting KREATE_DATA to default value => %+v", setErr)
		}
	}

	chartsLocation, ok = os.LookupEnv("KREATE_PROFILE")
	if !ok {
		setErr := os.Setenv("KREATE_PROFILE", PROFILES)
		if setErr != nil {
			log.Panicf("Error Setting KREATE_PROFILE to default value => %+v", setErr)
		}
	}
}

func InitializeEnvironment() {
	// TODO check if helm version is 2.16.3. If not, prompt the user to overwrite?
	// TODO check if istio environment is already deployed?
	runInstallScript()
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
