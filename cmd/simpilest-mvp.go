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
		command := exec.Command("sh", "./setup1.sh")
		command.Stderr = os.Stderr
		out, err := command.Output()
		if err != nil {
			log.Printf("Command Failed :: %s", err)
		}
		fmt.Print(string(out))
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("Please wait. Still deploying services.")
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("Please wait. Still deploying services.")
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("Please wait. Still deploying services.")
		command1 := exec.Command("sh", "./setup2.sh")
		command1.Stderr = os.Stderr
		out1, err1 := command1.Output()
		if err1 != nil {
			log.Printf("Command Failed :: %s", err1)
		}
		fmt.Print(string(out1))

	}

}

func errorHandler(err error) {
	if err != nil {
		log.Printf("Command Failed :: %s\n", err)
	}
}
