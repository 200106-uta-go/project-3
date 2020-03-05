package kreate

import (
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

/*
## kreate chart <profile name>
1. Part 1. (to be reused within kreate run)
    - The specified profile must be converted from its profile format to the helm chart's values.yaml format
    - The created values.yaml must then be copied to the kustom chart.
2. Part 2.
    - The kustom chart must then be copied to a user-friendly directory (implementation as descretioned by developer)
*/

//CreateChart creates a helm chart using the data provided in profile
// TODO -Need to add full path to file name to get correct yaml.
func CreateChart(profileName string) {
	profile := GetProfile(profileName + ".yaml")

	createValues(profile)
	createChartFile(profile)

	//build file structure for running helm
	buildFileSystem()

	//add values into chart for deployment yaml
	populateChart("values.yaml", "./templates")
}

//createValues creates a values.yaml based on a profile
func createValues(profile Profile) {
	//create values yaml
	file, err := os.Create("values.yaml")
	if err != nil {
		panic(err)
	}

	bytes, err := yaml.Marshal(profile)
	if err != nil {
		panic(err)
	}

	written, err := file.Write(bytes)
	if written == 0 {
		panic("Nothing was written to values.yaml")
	}
	if err != nil {
		panic(err)
	}
}

//populateChart injects the values inside filename into a chart template
func populateChart(filename string, templateDir string) {
	//uses helm to inject values into template

	cmd := exec.Command("helm", "template", "--output-dir", "./", "./")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//createChartFile generates the required chart.yaml metadata file to use with helm
func createChartFile(profile Profile) {
	chart := fmt.Sprintf(`apiVersion: v1
name: %s
version: 1.0.0
description: A custom ingress controller to provide failover requests to another address
keywords:
- ingress
- failover
sources:
- https://github.com/200106-uta-go/project-3
maintainers:
- name: do we want our names here? for posterity/blame`, profile.Name)

	chartFile, err := os.Create("Chart.yaml")
	if err != nil {
		panic(err)
	}

	chartFile.WriteString(chart)
}

//buildFileSystem sets up the file structure to install and template a helm chart
func buildFileSystem() {
	if !dirExists("./templates") {
		os.Mkdir("templates", 0777)
	}

	//copy files from /var/local/kreate into ./templates
	cmd := exec.Command("sudo", "cp", "-r", MOULDFOLDERS, "./templates")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//dirExists returns a boolean indicating if the given directory exists
func dirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
