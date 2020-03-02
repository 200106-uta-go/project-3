package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/200106-uta-go/project-3/pkg/executil"
)

const (
	script1 = `
	sudo curl -L https://istio.io/downloadIstio | sh -
	cd istio-1.4.5
	export PATH=$PWD/bin:$PATH
	sudo curl -L https://get.helm.sh/helm-v2.16.3-linux-amd64.tar.gz -o helm.tar.gz
	tar xf helm.tar.gz
	cd linux-amd64/
	sudo cp helm /bin/helm
	sudo cp tiller /bin/tiller
	cd ..
	sudo kubectl apply -f install/kubernetes/helm/helm-service-account.yaml
	sudo helm init --service-account tiller
	`
	script2 = ``
)

var args []string

func init() {
	flag.Parse()
	args = flag.Args()
}

func main() {
	//Which istio install package do you want?
	//demo setup
	var again string
	fmt.Println("=======================KREATE=========================")
	fmt.Println("Deploy istio, grafana, jaeger into cluster? y/n")
	fmt.Scan(&again)
	if again == "y" || again == "Y" || again == "Yes" || again == "yes" {

		fmt.Println("Deploying services to cluster. This may take upto few minutes.")
		//command1 := exec.Command("sh", "./setup1.sh")
		//command1.Stderr = os.Stderr
		//out1, err1 := command1.Output()
		_, err := executil.ExecHandler(script1, 0)
		errorHandler(err)
		//fmt.Print(string(out1))
		for i := 1; i <= 3; i++ {
			fmt.Println("Please wait. Still deploying services.")
			time.Sleep(time.Duration(10) * time.Second)
		}
		command2 := exec.Command("sh", "./setup2.sh")
		command2.Stderr = os.Stderr
		out2, err2 := command2.Output()
		errorHandler(err2)
		fmt.Print(string(out2))

	}

}

func errorHandler(err error) {
	if err != nil {
		log.Printf("Command Failed :: %s\n", err)
	}
}
