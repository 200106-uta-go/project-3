package kubeutil

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

// TearDown will first drain and then delete all pods on slave nodes.
func TearDown() error {
	// kubectl scale deployment.v1.apps/collider-deployment --replicas 0
	podCount := "0"
	out, err := exec.Command("sudo", "kubectl", "scale", "deployment.v1.apps/collider-deployment", "--replicas", podCount).Output()
	time.Sleep(time.Duration(20) * time.Second)
	if err != nil {
		return fmt.Errorf("Could not kubectl delete: Error==%+v", err)
	}

	log.Printf("-Teardown Successful->%s\n", out)

	/* <--------------------------------Drain and Delete Nodes-------------------------------->
	These are the command we are using to drain and then delete the nodes.
	sudo kubectl drain ip-172-31-32-93 --ignore-daemonsets --delete-local-data
	sudo kubectl delete node <node-name>

	Be sure to check node presence in a loop before leaving the TearDown() Function
	*/

	fmt.Println("=================START DRAINING NODE PHASE=======================")
	myNodes := RemoveMasterNode(NodeInfo())
	errNodes := make([]Node, 0)

	// Loop over nodes and drain them. If they return an error, add them to errNode list for error reporting.
	for i, v := range myNodes {
		out, err := exec.Command("sudo", "kubectl", "drain", v.Name, "--ignore-daemonsets", "--delete-local-data").Output()
		fmt.Printf("Draining Node [%d at %s] -> %s\n", i, v.Name, out)
		if err != nil {
			fmt.Printf("Error draining Node [%d at %s] -> %+v\n", i, v.Name, err)
			errNodes = append(errNodes, v)
		}
	}

	// Break out and return errNode list if unsuccessful.
	if len(errNodes) > 0 {
		fmt.Println("!NODE DRAINING UNSUCCESSFUL!")
		for _, v := range errNodes {
			fmt.Printf("Node %s not drained!\n", v.Name)
		}
		fmt.Println("Cancelling Node Deletion.")
		return fmt.Errorf("Nodes not deleted => %+v", errNodes)
	}

	// Wait for Nodes to Properly drain before moving on to deletion phase.
	isUpdated := false //When this is true break out of checking loop
	const drainCheck = "NotReady,SchedulingDisabled"

	for isUpdated {
		mapNodes := MapNodes(RemoveMasterNode(NodeInfo()))
		for k, v := range mapNodes {
			if v.Status == drainCheck {
				log.Printf("Node %s drained.\n", k)
				delete(mapNodes, k)
			} else {
				fmt.Printf("Node %s is still draining...\n", k)
			}
		}
		// Have all nodes that have finished draining been removed from the map?
		// If so, break from loop and move onto the deletion phase.
		if len(mapNodes) == 0 {
			isUpdated = true
		}
	}

	fmt.Println("=================END DRAINING NODE PHASE ==========================")
	fmt.Println("=================START DELETING NODE PHASE=========================")
	cantDelete := make([]Node, 0)

	// Delete Nodes now that they are drained.
	// Loop over nodes and drain them. If they return an error, add them to errNode list for error reporting.
	for i, v := range myNodes {
		out, err := exec.Command("sudo", "kubectl", "delete", "node", v.Name).Output()
		fmt.Printf("Deleting Node [%d at %s] -> %s\n", i, v.Name, out)
		if err != nil {
			fmt.Printf("Error deleting Node [%d at %s] -> %+v\n", i, v.Name, err)
			cantDelete = append(cantDelete, v)
		}
	}

	// Break out and return errNode list if unsuccessful.
	if len(cantDelete) > 0 {
		fmt.Println("!NODE DELETION UNSUCCESSFUL!")
		for _, v := range cantDelete {
			fmt.Printf("Node %s not deleted!\n", v.Name)
		}
		fmt.Println("Will return undeleted nodes.")
		return fmt.Errorf("Nodes not deleted => %+v", errNodes)
	}

	fmt.Println("=================END DELETING NODE PHASE===========================")
	fmt.Println("=================START DELETING SERVICE PHASE=========================")
	// Start Process of Deleteing Unneeded Services
	serviceOut, serviceErr := exec.Command("sudo", "kubectl", "delete", "service", "collider-service").Output()
	fmt.Printf("Deleting Service [%s] -> %s\n", "collider-service", serviceOut)
	if serviceErr != nil {
		fmt.Printf("Error deleting Service [%s] -> %+v\n", "collider-service", serviceErr)
	}

	fmt.Println("=================END DELETING SERVICE PHASE===========================")

	// Return Nil if we have success in all stages
	return nil
}

// MapNodes turns a slice of nodes into a map with each key == external ip
func MapNodes(nodes []Node) map[string]Node {
	mNodes := make(map[string]Node, 0)

	for _, v := range nodes {
		mNodes[v.Name] = v
	}

	return mNodes
}

// RemoveMasterNode takes the master node out of a slice if it is present.
// This will be useful for not acccidently deleting the master node during the
// tear down process.
func RemoveMasterNode(nodes []Node) []Node {
	// Finder Master Node Index
	var masterIndex int
	for i, v := range nodes {
		if v.Role == "master" {
			masterIndex = i
		}
	}

	// remove master from slice
	nodes[len(nodes)-1], nodes[masterIndex] = nodes[masterIndex], nodes[len(nodes)-1]
	return nodes[:len(nodes)-1]
}
