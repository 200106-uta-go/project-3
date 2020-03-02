package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/200106-uta-go/project-3/pkg/yaml/structgen"
)

func main() {
	file, err := os.Open("pkg/yaml/masterDeployment.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	gen := structgen.FromFile(file)
	spec := reflect.ValueOf(gen["spec"])
	fmt.Println(spec)
	fmt.Println(gen.GetKey("cpu"))
}
