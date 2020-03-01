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
	command := exec.Command("helm", args...)
	command.Stderr = os.Stderr
	out, err := command.Output()
	if err != nil {
		log.Printf("Command Failed :: %s", err)
	}
	fmt.Print(string(out))
}
