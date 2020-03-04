package gen

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// PortalGen creates yaml file for portal
func PortalGen() {
	var portal, targetIP string

	fmt.Printf("Enter name of Portal: ")
	fmt.Scanln(&portal)

	fmt.Printf("Enter target IP: ")
	fmt.Scanln(&targetIP)

	file, err := os.Create(portal + "Portal.yaml")
	errorHandler(err)
	defer file.Close()

	//marshal user data into ingress struct
	bytes, err := yaml.Marshal(dataWrite(portal, targetIP))
	errorHandler(err)

	_, err = file.Write(bytes)
	errorHandler(err)
}

// Add default values to Portal struct
func dataWrite(portal, target string) Portal {
	var data Portal

	data.APIVersion = "revature.com/v1"
	data.Kind = "Portal"
	data.Metadata.Name = "portal-" + portal
	data.Spec.Portal = portal
	data.Spec.TargetIP = target

	return data
}

// Generic error handler
func errorHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
