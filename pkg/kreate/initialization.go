package kreate

import (
	"fmt"
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

	// testing
	// InitHelm()
	// InitIstio()
	// InitKreate()
	// RemoveArtifacts()
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
	sudo kubectl apply -f install/kubernetes/helm/helm-service-account.yaml
	sudo helm init --service-account tiller
	echo "waiting for tiller pod to be ready ..."
	sudo kubectl -n kube-system wait --for=condition=Ready pod -l name=tiller --timeout=300s
	sudo helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
	echo "Waiting for istio-system jobs to complete (this may take a moment)..."
	kubectl -n istio-system wait --for=condition=complete job --all
	sudo helm install install/kubernetes/helm/istio --name istio --namespace istio-system --values install/kubernetes/helm/istio/values-istio-demo.yaml
	kubectl label namespace default istio-injection=enabled
	kubectl apply -f ../../../deployments/istio_env/istiometrics.yaml` // script goes here

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

	//Download helm.tar
	fmt.Println("LOG: Downloading Helm...")
	runcmd("sudo curl -L https://get.helm.sh/helm-v2.16.3-linux-amd64.tar.gz -o helm.tar.gz", home)

	//Extract Helm
	fmt.Println("LOG: Extracting Helm...")
	runcmd("tar -xf helm.tar.gz", home)

	//Copy helm and tiller executable from home/linux-amd64/ into the bin folder
	fmt.Println("LOG: Setting Helm...")
	runcmd("sudo cp helm /bin/helm", home+"/linux-amd64/")
	runcmd("sudo cp tiller /bin/tiller", home+"/linux-amd64/")
}

func InitIstio() {
	home, _ := os.UserHomeDir()

	//Downloading Istio 1.4.5 to user's home dir
	fmt.Println("LOG: Downloading Istio...")
	runcmd("sudo curl -L https://github.com/istio/istio/releases/download/1.4.5/istio-1.4.5-linux.tar.gz -o istio-1.4.5.tar.gz", home)

	//Extract Istio
	fmt.Println("LOG: Extracting Istio...")
	runcmd("tar -xf istio-1.4.5.tar.gz", home)

	//Move istioctl into /bin (can still use istioctl if wanted)
	fmt.Println("LOG: Setting Istioctl...")
	runcmd("sudo cp istio-1.4.5/bin/istioctl /bin/istioctl", home)
}

func InitKreate() {
	home, _ := os.UserHomeDir()

	// sudo apt install -y subversion
	fmt.Println("LOG: Installing Subversion (to get configs)...")
	runcmd("sudo apt install -y subversion", home)

	// svn export http://github.com/200106-uta-go/project-3/trunk/deployments/templates
	runcmd("svn export http://github.com/200106-uta-go/project-3/trunk/deployment/templates", home)

	// sudo mv templates/*.yaml  /var/local/kreate/
	runcmd("sudo mv templates/*.yaml "+MOULDFOLDERS, home)

	// get custome promethus config map
	// cp ../configmap.yaml install/kubernetes/helm/istio/charts/prometheus/templates/configmap.yaml -f

	// ASSUMES InitHelm() has already ran
	// install/inti tiller on cluster
	fmt.Println("LOG: Installing Helm Tiller...")
	runcmd("sudo kubectl apply -f istio-1.4.5/install/kubernates/helm-servive-account.yaml", home)
	runcmd("sudo helm init --service-account tiller", home)

	// wait for tiller to deploy
	fmt.Println("LOG: Waiting for tiller pod to be ready ...")
	//used cmd.Run because it needs to block execution
	cmd := exec.Command("/bin/sh", "-c", "kubectl -n default wait --for condition=established --timeout=60s crd/portals.revature.com")
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd.Stderr)
		// panic(err)
	}
	//runcmd("sudo kubectl -n kube-system wait --for=condition=Ready pod -l name=tiller --timeout=300s", home)

	// ASSUMES InitIstio() has already ran
	// init istio on kluster
	fmt.Println("LOG: Initilize Istio CRDs to cluster...")
	runcmd("sudo helm install istio-1.4.5/install/kubernetes/helm/istio-init --name istio-init --namespace istio-system", home)

	// wait for istio to be ready
	fmt.Println("LOG: Waiting for istio-system jobs to complete (this may take a bit)")
	//used cmd.Run because it needs to block execution
	cmd = exec.Command("/bin/sh", "-c", "kubectl -n default wait --for condition=established --timeout=60s crd/portals.revature.com")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(cmd.Stderr)
		// panic(err)
	}
	//runcmd("kubectl -n istio-system wait --for=condition=complete job --all", home)

	// install istio to the cluster
	fmt.Println("LOG: Deploy Istio (demo) to cluster...")
	runcmd("sudo helm install istio-1.4.5/install/kubernetes/helm/istio --name istio --namespace istio-system --values istio-1.4.5/install/kubernetes/helm/istio/values-istio-demo.yaml", home)

	// allow istio sidecar injection on default namespace
	fmt.Println("LOG: Enabling Istio side-car injection...")
	runcmd("kubectl label namespace default istio-injection=enabled", home)

	// expose telemetrics (Prometheus -port:15030 / Grafan -port:15031 / Jaeger -port:15032)
	// need to make this a constent from somewhere online (repo)
	fmt.Println("LOG: Exposing Istio telemetrics (Prometheus -port:15030 / Grafan -port:15031 / Jaeger -port:15032)...")
	runcmd("kubectl apply -f deployments/istio_env/istiometrics.yaml", home+"/go/src/github.com/200106-uta-go/project-3/")
}

func RemoveArtifacts() {
	home, _ := os.UserHomeDir()

	// remove helm's download
	runcmd("rm -R linux-amd64", home)
	runcmd("rm helm.tar.gz", home)

	// remove istio's dowloads
	// currently leaving istio-1.4.5 folder till verified not needed after init
	// runcmd("rm -R istio-1.4.5", home)
	runcmd("istio-1.4.5.tar.gz", home)
}

//Runs Com Command in Dir Directory
func runcmd(com string, dir string) {
	cmd := exec.Command("sh", "-c", com)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(out))
}
