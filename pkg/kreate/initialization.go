package kreate

import (
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

func Initalization() { // current logic was written prior to the 3/3/20 MVP meeting

	home, _ := os.UserHomeDir()

	pathErr := os.MkdirAll(MOULDFOLDERS, 1777)
	if pathErr != nil {
		log.Panicf("Error making directory %s => %v", MOULDFOLDERS, pathErr)
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
