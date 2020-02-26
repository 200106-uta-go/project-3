//Scanner pulls information from the kubernetes cluster that is
//running locally on the machine. It does this every TIMETOSLEEP
//in a constant loop. It stores this information in various files.
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

//TIMETOSLEEP this is the amount of time that the program waits
//between resource scans.
const TIMETOSLEEP = 10 * time.Second

//Node contains all used information of what it consists of.
type Node struct {
	Name        string
	Status      string
	Roles       string
	Age         string
	Version     string
	Description string
}

//Pod contains all used information of what it consists of
type Pod struct {
	Name        string
	Ready       string
	Status      string
	Restarts    string
	Age         string
	Port        string
	Description string
}

//Service contains all used information of what it consists of
type Service struct {
	Name        string
	Type        string
	ClusterIP   string
	ExternalIP  string
	Port        string
	Age         string
	Description string
}

//Deployment contains all used information of what it consists of
type Deployment struct {
	Name        string
	Ready       string
	UpToDate    string
	Available   string
	Age         string
	Description string
}

func main() {
	go GetNodes()
	go GetPods()
	go GetDeployments()
	GetServices()
}

//GetNodes Scans for all nodes and serilizes them as a
//Node struct and saved to a json file.
func GetNodes() {
	for {
		var NewNode Node
		var TempNodesList []Node

		output, _ := exec.Command("kubectl", "get", "nodes").Output()

		t := strings.Split(string(output), "\n")
		t = t[1:]

		for _, v := range t {
			z := strings.Split(v, " ")

			var temp []string

			for k2, v2 := range z {
				z[k2] = strings.TrimSpace(v2)
				if z[k2] != "" {
					temp = append(temp, z[k2])
				}
			}

			z = temp

			if len(z) != 0 {
				descrip, _ := exec.Command("kubectl", "describe", "nodes", z[0]).Output()

				NewNode = Node{Name: z[0], Status: z[1], Roles: z[2], Age: z[3], Version: z[4], Description: string(descrip)}
				TempNodesList = append(TempNodesList, NewNode)
			}

		}

		byteslice, _ := json.MarshalIndent(TempNodesList, "", "	")

		ioutil.WriteFile("../nodes.json", byteslice, 7777)

		time.Sleep(TIMETOSLEEP)
	}
}

//GetPods Scans for all pods and serilizes them as a
//Node struct and saved to a json file.
func GetPods() {
	for {
		var NewPod Pod
		var TempPodList []Pod

		// Get a list of current pods from k8s
		output, _ := exec.Command("kubectl", "get", "pods").Output()

		// Seperate each pods by splitting on new lines
		t := strings.Split(string(output), "\n")
		t = t[1:]

		// For each pod make a pod struct and add it to a slice of pod structs
		for _, v := range t {
			// Seperate each word
			z := strings.Split(v, " ")

			var temp []string

			// Trim spaces from words
			for k2, v2 := range z {
				z[k2] = strings.TrimSpace(v2)
				if z[k2] != "" {
					temp = append(temp, z[k2])
				}
			}

			// Set the slice of words to the trimmed slice of words
			z = temp

			// If there is a pod to be looked at
			if len(z) != 0 {
				// Get the decription of the pod
				descrip, _ := exec.Command("kubectl", "describe", "pods", z[0]).Output()
				descripfile, _ := os.OpenFile("./pods", os.O_RDWR|os.O_CREATE, 7777)
				descripfile.Write(descrip)

				grepport, _ := exec.Command("grep", "Port:", "./pods").Output()
				grepportslice := strings.Split(string(grepport), "\n")
				grepportslice = strings.Split(string(grepportslice[0]), " ")

				port := grepportslice[len(grepportslice)-1]
				portslice := strings.Split(port, "/")

				NewPod = Pod{Name: z[0], Ready: z[1], Status: z[2], Restarts: z[3], Age: z[4], Port: portslice[0], Description: string(descrip)}
				TempPodList = append(TempPodList, NewPod)
			}

		}

		byteslice, _ := json.MarshalIndent(TempPodList, "", "	")

		ioutil.WriteFile("../pods.json", byteslice, 7777)

		time.Sleep(TIMETOSLEEP)
	}
}

//GetServices Scans for all services and serilizes them as a
//Node struct and saved to a json file.
func GetServices() {
	for {
		var NewService Service
		var TempNewServiceList []Service

		// Get the services from k8s
		output, _ := exec.Command("kubectl", "get", "svc").Output()

		// Seperate services by the new line
		t := strings.Split(string(output), "\n")
		t = t[1:]

		// For each service populate a Service struct and add it to the
		// TempNewServiceList []Service
		for _, v := range t {
			// Seperate each word
			z := strings.Split(v, " ")

			// Slice to hold strings temporarily
			var temp []string

			// Trim spaces for each word
			for k2, v2 := range z {
				z[k2] = strings.TrimSpace(v2)
				if z[k2] != "" {
					temp = append(temp, z[k2])
				}
			}

			// reset z to the trimmed version temp
			z = temp

			// if there is a service
			if len(z) != 0 {
				// Get the description of the service from k8s
				descrip, _ := exec.Command("kubectl", "describe", "svc", z[0]).Output()

				// Make a Service struct with the information on the service from k8s
				NewService = Service{Name: z[0], Type: z[1], ClusterIP: z[2], ExternalIP: z[3], Port: z[4], Age: z[5], Description: string(descrip)}
				// Append the Service struct to the TempNewServiceList
				TempNewServiceList = append(TempNewServiceList, NewService)
			}

		}

		// Write the new list of services to the services.json
		byteslice, _ := json.MarshalIndent(TempNewServiceList, "", "	")

		ioutil.WriteFile("../services.json", byteslice, 7777)

		time.Sleep(TIMETOSLEEP)
	}
}

//GetDeployments Scans for all deployments and serilizes them as a
//Node struct and saved to a json file.
func GetDeployments() {
	// Continuously check
	for {
		var NewDeployment Deployment
		var TempNewDeploymentList []Deployment

		// Get the active deployments from the k8s cluster
		output, _ := exec.Command("kubectl", "get", "deployments").Output()

		// Seperate the deployments by splitting on new line
		t := strings.Split(string(output), "\n")
		t = t[1:]

		// For each deployment get the description, populate a Deployment struct and
		// add it to the TempNewDeploymentList
		for _, v := range t {
			// Split individual words
			z := strings.Split(v, " ")

			var temp []string

			// For each word in the line remove any extra spaces
			for k2, v2 := range z {
				z[k2] = strings.TrimSpace(v2) // remove spaces from the line
				if z[k2] != "" {              // Check if there is anything there
					temp = append(temp, z[k2]) // add the word to the new set of words
				}
			}

			// Set z equal to the trimmed set of words
			z = temp

			// If there is a deployment then get its description
			if len(z) != 0 {
				// Get description of the deployment
				descrip, _ := exec.Command("kubectl", "describe", "deployments", z[0]).Output()

				// Make a Deployment struct with the information returned from k8s
				NewDeployment = Deployment{Name: z[0], Ready: z[1], UpToDate: z[2], Available: z[3], Age: z[4], Description: string(descrip)}
				// Append the Deployment struct to the deploymentList
				TempNewDeploymentList = append(TempNewDeploymentList, NewDeployment)
			}

		}

		// Make a byte slice out of the deployment list
		byteslice, _ := json.MarshalIndent(TempNewDeploymentList, "", "	")

		// Write the new deployment to the deployments.json file
		ioutil.WriteFile("../deployments.json", byteslice, 7777)

		// Wait TIMETOSLEEP before checking again
		time.Sleep(TIMETOSLEEP)
	}
}
