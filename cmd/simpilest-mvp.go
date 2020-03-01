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
	fmt.Println("=======================KREATE=========================")
	fmt.Println("Deploy istio, grafana, jaeger into cluster? y/n")
	fmt.Scan(&yes)
	if again == "y" || again == "Y" || again == "Yes" || again == "yes" {

		command := exec.Command("./setup.sh")
		command.Stderr = os.Stderr
		out, err := command.Output()
		if err != nil {
			log.Printf("Command Failed :: %s", err)
		}
		fmt.Print(string(out))

	}

}
