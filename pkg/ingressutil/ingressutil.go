package ingressutil

/*
Ingressutil package created by Matt Ackard and Josh Nguyen.
For use in adding a new rule to a kubernetes ingress deployment
This package assumes you define the ingress rules in the top of your yaml file
*/

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//AddPath adds the ingress route to an existing ingress yaml file
func AddPath(file *os.File, host, path, svcName string, svcPort int) error {

	stringSlice, err := separateYAML(file)
	check(err)

	//creates a new ingress struct and unmarshalse ingress section into it
	ingress, err := addRule(stringSlice, host, path, svcName, svcPort)
	check(err)

	newBytes, err := yaml.Marshal(ingress)
	check(err)

	//replace first section of yaml to add new rules
	newBytes = append(newBytes, []byte("---")...)

	for i, section := range stringSlice[1:] {
		//don't add seapartor dashed on final yaml section
		if i == len(stringSlice[1:])-1 {
			newBytes = append(newBytes, []byte(section)...)
		} else {
			newBytes = append(newBytes, []byte(section+"---")...)
		}
	}

	err = ioutil.WriteFile(file.Name(), newBytes, 0644)
	check(err)

	return nil
}

//separateYAML splits multiple deployments in a single yaml file
func separateYAML(file *os.File) ([]string, error) {
	//open the file and read its contents
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	//convert contents to string and split on ---
	contents := string(bytes)
	stringSlice := strings.Split(contents, "---")
	return stringSlice, nil
}

//addRule appends a new rule to the ingress deployment
func addRule(stringSlice []string, host string, path string, svcName string, svcPort int) (Ingress, error) {
	//creates a new ingress struct and unmarshalse ingress section into it
	ingress := Ingress{}
	err := yaml.Unmarshal([]byte(stringSlice[0]), &ingress)
	if err != nil {
		return Ingress{}, err
	}

	//append new rile to the ingress struct
	ingress.Spec.Rules = append(ingress.Spec.Rules, Rule{
		host,
		HTTP{
			[]Paths{
				{
					path,
					Backend{
						svcName,
						svcPort,
					},
				},
			},
		},
	})

	return ingress, nil
}
