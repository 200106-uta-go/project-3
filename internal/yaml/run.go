package main

import (
	"fmt"
	"log"
	"os"

	"github.com/200106-uta-go/project-3/pkg/yaml/structgen"
	"github.com/200106-uta-go/project-3/pkg/yaml/yamlgen"
)

func main() {
	file, err := os.Open("pkg/yaml/masterDeployment.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	gen := structgen.FromFile(file)
	fmt.Println(gen.GetKey("cpu"))
	fmt.Print("\n\n")

	yamlgen.FromImage("mattackard/kube-portal:v6.0")
}
