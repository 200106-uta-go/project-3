package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const filename = "tempSetup.sh"

const script = `sudo curl -L https://istio.io/downloadIstio | sh -
cd istio-*
export PATH=$PWD/bin:$PATH
sudo curl -L https://get.helm.sh/helm-v2.16.3-linux-amd64.tar.gz -o helm.tar.gz
tar xf helm.tar.gz
cd linux-amd64/
sudo cp helm /bin/helm
sudo cp tiller /bin/tiller
cd ..
sudo kubectl apply -f install/kubernetes/helm/helm-service-account.yaml
sudo helm init --service-account tiller
echo "waiting for tiller pod to be ready ..."
sudo kubectl -n kube-system wait --for=condition=Ready pod -l name=tiller --timeout=300s
sudo helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
echo "waiting for istio-system jobs to complete (may take about a min)"
kubectl -n istio-system wait --for=condition=complete job --all
sudo helm install install/kubernetes/helm/istio --name istio --namespace istio-system --values install/kubernetes/helm/istio/values-istio-demo.yaml
kubectl create configmap ingress --from-file=${HOME}/.kube/config
kubectl apply -f portalCRD.yml` // script goes here

var args []string

func main() {
	//Which istio install package do you want?
	//demo setup
	var again string
	fmt.Println("=======================KREATE=========================")
	fmt.Println("Deploy istio, grafana, jaeger into cluster? y/n")
	fmt.Scan(&again)
	if again == "y" || again == "Y" || again == "Yes" || again == "yes" {

		er := ioutil.WriteFile(filename, []byte(script), 0777)
		if er != nil {
			log.Fatal("Script failed to deploy - ", er)
		}

		cmd := exec.Command("/bin/sh", filename)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		er = cmd.Run()
		if er != nil {
			log.Fatal("Script failed to deploy - ", er)
		}

		os.Remove(filename)
	}

}
