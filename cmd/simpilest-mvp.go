package main

import (
	"flag"
	"fmt"

	"github.com/200106-uta-go/project-3/pkg/executil"
)

const (
	script = `sudo curl -L https://istio.io/downloadIstio -o istio.sh
	ls -l
	sh istio.sh
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
	helm install install/kubernetes/helm/istio-init --name istio-init --namespace istio-system
	kubectl -n istio-system wait --for=condition=complete job --all
	helm install install/kubernetes/helm/istio --name istio --namespace istio-system --values install/kubernetes/helm/istio/values-istio-demo.yaml`

	script2 = `echo "Hello World"
	sleep 2
	echo "World Hello"`
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

		executil.ExecHandler(script)
		// lineNumber, err := executil.ExecHandler(script1, 0)
		// if err != nil {
		// 	log.Printf("Restarting Script at Line => %d", lineNumber)
		// 	newLIne, errNew := executil.ExecHandler(script1, lineNumber)
		// 	if err != nil {
		// 		log.Printf("Restarting script AGAIN at Line => %d\n Error => %v\n", newLIne, errNew)
		// 		newLine2, errNew2 := executil.ExecHandler(script1, newLIne)
		// 		if err != nil {
		// 			log.Printf("No more attemps. Line Number => %d\nError => %v", newLine2, errNew2)
		// 		}
		// 	}
		// }
		//fmt.Print(string(out1))

		// for i := 1; i <= 3; i++ {
		// 	fmt.Println("Please wait. Still deploying services.")
		// 	time.Sleep(time.Duration(10) * time.Second)
		// }
		// command2 := exec.Command("sh", "./setup2.sh")
		// command2.Stderr = os.Stderr
		// out2, err2 := command2.Output()
		// errorHandler(err2)
		// fmt.Print(string(out2))

		// lineNumberPart2, errPart2 := executil.ExecHandler(script2, 0)
		// if errPart2 != nil {
		// 	log.Printf("Restarting Script at Line => %d", lineNumberPart2)
		// 	newLIne, errNew := executil.ExecHandler(script2, lineNumberPart2)
		// 	if err != nil {
		// 		log.Printf("Restarting script AGAIN at Line => %d\n Error => %v\n", newLIne, errNew)
		// 		newLine2, errNew2 := executil.ExecHandler(script2, newLIne)
		// 		if err != nil {
		// 			log.Printf("No more attemps. Line Number => %d\nError => %v", newLine2, errNew2)
		// 		}
		// 	}
		// }

	}
}
