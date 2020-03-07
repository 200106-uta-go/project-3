package kreate

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

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
func CreateChart(profileName string) {

	if profileName == "" {
		fmt.Println("Profile name is required to create a chart")
		showChartHelp()
	} else {
		profile := GetProfile(profileName)

		//build file structure for running helm
		buildFileSystem(profile)

		createValues(profile)
		createChartFile(profile)

		//add values into chart for deployment yaml
		populateChart("values.yaml", "./charts/"+profile.Name)

		//update file permissions and reorganize directories
		fixFileSystem(profile)
	}
}

func showChartHelp() {
	help := `Usage:
	kreate chart [profile name]

Examples:
	kreate chart myProfile
	kreate chart anotherProfile.yaml`
	fmt.Print(help, "\n\n")
}

//createValues creates a values.yaml based on a profile
func createValues(profile Profile) {
	//create values yaml
	file, err := os.Create("./charts/" + profile.Name + "/values.yaml")
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

//populateChart injects the values inside valuesFile into chart templates
func populateChart(valuesFile string, chartDir string) {
	//uses helm to inject values into template
	if !strings.HasSuffix(valuesFile, ".yaml") {
		valuesFile += ".yaml"
	}
	if !dirExists(chartDir + "/deploy") {
		err := os.Mkdir(chartDir+"/deploy", 0777)
		if err != nil {
			panic(err)
		}
	}
	cmd := exec.Command("helm", "template", "--values", chartDir+"/"+valuesFile, "--output-dir", chartDir+"/deploy", chartDir)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(cmd.Stderr)
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

	chartFile, err := os.Create("./charts/" + profile.Name + "/Chart.yaml")
	if err != nil {
		panic(err)
	}

	chartFile.WriteString(chart)
}

//buildFileSystem sets up the file structure to install and template a helm chart
func buildFileSystem(profile Profile) {
	if !dirExists("./charts/" + profile.Name + "/templates/") {
		os.MkdirAll("./charts/"+profile.Name+"/templates/", 0777)
	}

	//copy files from /var/local/kreate into ./templates
	copyDir(MOULDFOLDERS, "./charts/"+profile.Name+"/templates/")
}

//fixFileSystem rearranges the generated files to better organize the directory and to give user permissions
func fixFileSystem(profile Profile) {
	//move templates
	copyDir("./charts/"+profile.Name+"/deploy/"+profile.Name+"/templates/", "./charts/"+profile.Name+"/deploy/")

	//delete empty files in deploy folder
	cmd2 := exec.Command("rm", "-r", "./charts/"+profile.Name+"/deploy/"+profile.Name)
	err := cmd2.Run()
	if err != nil {
		panic(err)
	}
}

//copyDir copies the contents of sourceDir into targetDir
func copyDir(sourceDir string, targetDir string) {
	//add slashes to directories if not already present
	if !strings.HasSuffix(sourceDir, "/") {
		sourceDir += "/"
	}
	if !strings.HasSuffix(targetDir, "/") {
		targetDir += "/"
	}

	//get a list of files/folders in source directory
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	//copy each file in source files to target directory
	for _, file := range files {
		cmd := exec.Command("cp", "-r", sourceDir+file.Name(), targetDir)
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			panic(cmd.Stderr)
		}
	}
}

//dirExists returns a boolean indicating if the given directory exists
func dirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
