package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const filename = "tempSetup.sh"

const script = `echo HELLO` // script goes here

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
