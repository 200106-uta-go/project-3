package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
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

		command := exec.Command("sh", "./setup.sh")
		command.Stderr = os.Stderr
		out, err := command.Output()
		if err != nil {
			log.Printf("Command Failed :: %s", err)
		}
		fmt.Print(string(out))
		command1 := exec.Command("sh", "./setup.sh")
		command1.Stderr = os.Stderr
		out1, err1 := command1.Output()
		if err1 != nil {
			log.Printf("Command Failed :: %s", err1)
		}
		fmt.Print(string(out1))

	}

}
