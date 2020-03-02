package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
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
		command1 := exec.Command("sh", "./setup1.sh")
		command1.Stderr = os.Stderr
		out1, err1 := command1.Output()
		errorHandler(err1)
		fmt.Print(string(out1))
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("Please wait. Still deploying services.")
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("Please wait. Still deploying services.")
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("Please wait. Still deploying services.")
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
