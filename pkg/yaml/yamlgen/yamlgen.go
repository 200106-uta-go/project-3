package yamlgen

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"strings"
)

//DockerInspect holds relevant data retrieved from the docker image inspect command
type DockerInspect struct {
	Config   DockerConfig `json:"Config"`
	RepoTags []string     `json:"RepoTags"`
	Name string
}

//DockerConfig holds configuration data for a docker imagef
type DockerConfig struct {
	ExposedPorts map[string]interface{}
}

//FromImage generates a yaml configuration from a docker image
func FromImage(image string) {
	//install docker if not on the system
	if !isInstalled("docker") {
		dockerInstall()
	}

	//pull docker image
	dockerPull(image)

	//inspect docker image and get exposed ports
	config := dockerInspect(image)

	//generate yaml exposing ports defined in docker image
	generateYAML(config)

	//remove image from system after done
	// dockerRemove(image)
}

//isInstalled checks to see if the command is installed on the system
func isInstalled(command string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command", "-v", command)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

//dockerInstall installs docker on the host machine if it doesn't already exist
func dockerInstall() {
	//use curl to get install script
	curl := exec.Command("curl", "-fsSL", "https://get.docker.com", "-o", "get-docker.sh")
	print("1")
	curl.Stdout = os.Stdout
	curl.Stderr = os.Stderr
	err := curl.Run()
	if err != nil {
		log.Fatalln(err)
	}

	//run script retrieved in curl request
	sh := exec.Command("sh", "get-docker.sh")
	print("2")
	sh.Stdout = os.Stdout
	sh.Stderr = os.Stderr
	err = sh.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

//dockerPull pulls the image from dockerhub to the host machine
func dockerPull(image string) {
	cmd := exec.Command("docker", "pull", image)
	print("3")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

//dockerInspect returns a struct containing the json values retrieved
//when running docker image inspect
func dockerInspect(image string) DockerInspect {
	cmd := exec.Command("docker", "image", "inspect", image)
	res, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	inspect := DockerInspect{}
	err = json.Unmarshal(res[1:len(res)-2], &inspect)
	if err != nil {
		log.Fatalln(err)
	}
	return inspect
}

func dockerRemove(image string) {
	cmd := exec.Command("docker", "image", "rm", image)
	print("5")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

//generateYAML creates a yaml deployment using exposed ports from the docker image
//and name/labels based on the image tags/name
func generateYAML(config DockerInspect) {
	var ports []string
	fmt.Println("Open ports: ")
	for key := range config.Config.ExposedPorts {
		fmt.Println(key)
		tempKey := key
		if (strings.Contains(tempKey, "/")) {
			tempKey = strings.Split(tempKey, "/")[0]
		}
		ports = append(ports, tempKey)
	}
	fmt.Println("Docker Image tags: ")
	for _, tag := range config.RepoTags {
		fmt.Println(tag)
	}

	//import the yaml template
	name := trimName(config.RepoTags[0])
	tmpl := template.Must(template.New(name).Parse(deployment))
	
	//create a struct to hold values for template
	values := struct{
		Name string
		Image string
		ExposedPorts []string
	}{
		Name: name,
		Image: config.RepoTags[0],
		ExposedPorts: ports,
	}

	//create a new yaml and populate it with config values
	newYAML, err := os.Create(name + ".yaml")
	err = tmpl.Execute(newYAML, values)
	if err != nil {
		log.Println("executing template:", err)
	}
}

//trimName takes a docker image name and trims anything after a / and before a :
func trimName(fullName string) string {
	prettyName := fullName
	if strings.Contains(prettyName, "/") {
		prettyName = strings.Split(prettyName, "/")[1]
	}
	if strings.Contains(prettyName, ":") {
		prettyName = strings.Split(prettyName, ":")[0]
	}
	return prettyName
}
