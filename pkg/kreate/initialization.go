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

	const installScript = `curl -L https://github.com/istio/istio/releases/download/1.4.5/istio-1.4.5-linux.tar.gz -o istio-1.4.5.tar.gz
	sudo apt install -y subversion 
	tar -xf istio-1.4.5.tar.gz
	cd istio-1.4.5
	export PATH=$PWD/bin:$PATH
	curl -L https://get.helm.sh/helm-v2.16.3-linux-amd64.tar.gz -o helm.tar.gz
	tar -xf helm.tar.gz
	cd linux-amd64/
	sudo cp helm /bin/helm
	sudo cp tiller /bin/tiller
	cd ..
	svn export http://github.com/200106-uta-go/project-3/trunk/deployments/templates
	sudo mv templates/*.yaml  /var/local/kreate/
	rm -r linux-amd64
	rm helm.tar.gz
	rm ../istio-1.4.5.tar.gz
	sudo chmod 777 ../istio-1.4.5
	# cp ../configmap.yaml install/kubernetes/helm/istio/charts/prometheus/templates/configmap.yaml -f
	sudo kubectl apply -f install/kubernetes/helm/helm-service-account.yaml
	sudo helm init --service-account tiller
	echo "waiting for tiller pod to be ready ..."
	sudo kubectl -n kube-system wait --for=condition=Ready pod -l name=tiller --timeout=300s
	sudo helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
	echo "waiting for istio-system jobs to complete (will take about minute)"
	kubectl -n istio-system wait --for=condition=complete job --all
	sudo helm install install/kubernetes/helm/istio --name istio --namespace istio-system --values install/kubernetes/helm/istio/values-istio-demo.yaml
	echo "Brandon locker's edits"
	cd ../../..
	echo $PWD
	kubectl label namespace default istio-injection=enabled
	kubectl apply -f deployments/terraform/dev_env/istio_env/istiometrics.yaml` // script goes here

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

func InitHelm() { 

	home, _ := os.UserHomeDir()

	workingDir, _ := os.Getwd()

	//Downloading Istio 1.4.5
	runcmd("sudo curl -L https://github.com/istio/istio/releases/download/1.4.5/istio-1.4.5-linux.tar.gz -o istio-1.4.5.tar.gz", home)

	//Extract Istio
	runcmd("tar -xf istio-1.4.5.tar.gz", home)

	//Move istioctl into /bin
	runcmd("sudo cp istio-1.4.5/bin/istioctl /bin/istioctl", home)

	//Download helm.tar
	runcmd("sudo curl -L https://get.helm.sh/helm-v2.16.3-linux-amd64.tar.gz -o helm.tar.gz", home)

	//Extract Helm
	runcmd("tar -xf helm.tar.gz", home)

	//Copy helm and tiller executable from home/linux-amd64/ into the bin folder
	runcmd("sudo cp helm /bin/helm", home+"/linux-amd64/")
	runcmd("sudo cp tiller /bin/tiller", home+"/linux-amd64/")
}

func InitIstio() { 

	home, _ := os.UserHomeDir()

	//Downloading Istio 1.4.5
	runcmd("sudo curl -L https://github.com/istio/istio/releases/download/1.4.5/istio-1.4.5-linux.tar.gz -o istio-1.4.5.tar.gz", home)

	//Extract Istio
	runcmd("tar -xf istio-1.4.5.tar.gz", home)

	//Move istioctl into /bin
	runcmd("sudo cp istio-1.4.5/bin/istioctl /bin/istioctl", home)

	//Download helm.tar
	runcmd("sudo curl -L https://get.helm.sh/helm-v2.16.3-linux-amd64.tar.gz -o helm.tar.gz", home)

	//Extract Helm
	runcmd("tar -xf helm.tar.gz", home)

	//Copy helm and tiller executable from home/linux-amd64/ into the bin folder
	runcmd("sudo cp helm /bin/helm", home+"/linux-amd64/")
	runcmd("sudo cp tiller /bin/tiller", home+"/linux-amd64/")
}

//Runs Com Command in Dir Directory
func runcmd(com string, dir string) {
	cmd := exec.Command("sh", "-c", com)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
