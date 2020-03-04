package ingress

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

//CreateIngress takes in ingress values and creates a kubernetes ingress yaml
func CreateIngress(filename string, ingress Ingress) {
	//add yaml extension if filename doesn't already have it
	if !strings.HasSuffix(filename, ".yaml") {
		filename += ".yaml"
	}

	//create the yaml file
	file, err := os.Create(filename)
	genericErrHandler(err)
	defer file.Close()

	ingress = applyIngressDefaults(ingress)
	ingress = validateIngress(ingress)

	//marshal user data into ingress struct
	bytes, err := yaml.Marshal(ingress)
	genericErrHandler(err)

	_, err = file.Write(bytes)
	genericErrHandler(err)
}

//adds default values to ingress in case user did not provide all values
func applyIngressDefaults(ingress Ingress) Ingress {
	//make sure apiVersion isn't old (extensions/v1beta1)
	ingress.APIVersion = "networking.k8s.io/v1beta1"

	//makes sure kind is always ingress
	if ingress.Kind != "Ingress" {
		if ingress.Kind != "" {
			fmt.Printf("Kind was set to %s ... Forcing kind Ingress \n", ingress.Kind)
		}
		ingress.Kind = "Ingress"
	}

	if ingress.Metadata.Name == "" {
		ingress.Metadata.Name = "default"
	}

	//if no backend or rule are defined, add a default backend to kubernetes default service
	if ingress.Spec.Backend == (Backend{}) && len(ingress.Spec.Rules) == 0 && len(ingress.Spec.TLS) == 0 {
		ingress.Spec.Backend = Backend{
			ServiceName: "kubernetes",
			ServicePort: 8443,
		}
	}

	return ingress
}

//checks for validity of picky ingress fields
func validateIngress(ingress Ingress) Ingress {
	// match, err := regexp.MatchString(`'[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')`, ingress.Metadata.Name)
	// genericErrHandler(err)
	// if match {
	// 	fmt.Println("Invalid name for ingress ... conforming name to requirements")
	// }
	return ingress
}

func genericErrHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
