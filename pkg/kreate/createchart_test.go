package kreate

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

var testProfile Profile

func TestMain(m *testing.M) {
	//create a test folder to hold temporary test files
	os.Mkdir("./chartTest", 0777)
	os.Chdir("chartTest")
	CreateProfile("test")

	//run through tests
	code := m.Run()

	//delete temporary test directory
	os.RemoveAll("../chartTest")

	//the end
	os.Exit(code)
}

func TestBuildFileSystem(t *testing.T) {
	//build file system
	buildFileSystem(testProfile)

	//check that file structure is accurate
	files, err := ioutil.ReadDir("./charts/" + testProfile.Name)
	if err != nil {
		t.Error(err)
	}

	if len(files) != 1 {
		filenames := ""
		for _, v := range files {
			filenames += v.Name() + ", "
		}
		t.Errorf("Filesystem expected 1 items, got %d: %s", len(files), filenames)
	}
	if files[0].Name() != "templates" {
		t.Error("Templates folder not created")
	}
}

func TestCreateValues(t *testing.T) {
	//build test charts dir
	os.MkdirAll("./charts/"+testProfile.Name, 0777)

	//generate values.yaml and read it
	createValues(testProfile)
	values, err := os.Open("./charts/" + testProfile.Name + "/values.yaml")
	if err != nil {
		t.Error(err)
	}
	valueContents, err := ioutil.ReadAll(values)
	if err != nil {
		t.Error(err)
	}

	//marshal values into a profile struct
	newProfile := Profile{}
	yaml.Unmarshal(valueContents, &newProfile)

	//compare the testProfile with the values profile
	if len(testProfile.Apps) != len(newProfile.Apps) {
		t.Errorf("App slice lengths do not match: profile %d | values %d", len(testProfile.Apps), len(newProfile.Apps))
	}
	if testProfile.ClusterIP != newProfile.ClusterIP {
		t.Errorf("IPs do not match: profile %s | values %s", testProfile.ClusterIP, newProfile.ClusterIP)
	}
	if testProfile.ClusterName != newProfile.ClusterName {
		t.Errorf("ClusterNames do not match: profile %s | values %s", testProfile.ClusterName, newProfile.ClusterName)
	}
	if testProfile.Name != newProfile.Name {
		t.Errorf("Names do not match: profile %s | values %s", testProfile.Name, newProfile.Name)
	}
	if len(testProfile.ClusterPorts) != len(newProfile.ClusterPorts) {
		t.Errorf("App slice lengths do not match: profile %d | values %d", len(testProfile.ClusterPorts), len(newProfile.ClusterPorts))
	}
}

func TestCreateChartFile(t *testing.T) {
	//build test charts dir
	os.MkdirAll("./charts/"+testProfile.Name, 0777)

	//copy of default Chart.yaml
	defaultChart := fmt.Sprintf(`apiVersion: v1
name: %s
version: 1.0.0
description: A custom ingress controller to provide failover requests to another address
keywords:
- ingress
- failover
sources:
- https://github.com/200106-uta-go/project-3
maintainers:
- name: do we want our names here? for posterity/blame`, testProfile.Name)

	//create the chart file
	createChartFile(testProfile)

	chart, err := os.Open("./charts/" + testProfile.Name + "/Chart.yaml")
	if err != nil {
		t.Error(err)
	}

	chartContents, err := ioutil.ReadAll(chart)
	if err != nil {
		t.Error(err)
	}

	if defaultChart != string(chartContents) {
		t.Error("Chart.yaml contents do not match default contents")
	}

}

func TestPopulateChart(t *testing.T) {
	const testValueYaml = `name: Test
clustername: TestCluster
clusterip: 127.0.0.1
clusterports:
- 8080
- 9090
app:
- name: TestApp
  imageurl: testapp.com/image
  ports: 
  - 8000
  - 9000
  endpoints: 
  - /one
  - /two
  - /three`

	//copy of default Chart.yaml
	const defaultChart = `apiVersion: v1
name: Test
version: 1.0.0
description: A custom ingress controller to provide failover requests to another address
keywords:
- ingress
- failover
sources:
- https://github.com/200106-uta-go/project-3
maintainers:
- name: do we want our names here? for posterity/blame`

	//build test charts dir
	os.MkdirAll("./charts/"+testProfile.Name+"/deploy/templates", 0777)

	//write test values.yaml
	er := ioutil.WriteFile("./charts/values.yaml", []byte(testValueYaml), 0777)
	if er != nil {
		t.Error("Values.yaml not created - ", er)
	}
	er = ioutil.WriteFile("./charts/Chart.yaml", []byte(defaultChart), 0777)
	if er != nil {
		t.Error("Chart.yaml not created - ", er)
	}

	populateChart("values.yaml", "./charts/"+testProfile.Name)

	//check that all deployments were created in charts/test/deploy/
	deploy, err := ioutil.ReadDir("./charts/" + testProfile.Name + "/deploy/" + testProfile.Name + "/templates")
	if err != nil {
		t.Error(err)
	}
	templates, err := ioutil.ReadDir("/var/local/kreate")
	if err != nil {
		t.Error(err)
	}
	for index, file := range templates {
		if file.Name() != deploy[index].Name() {
			t.Errorf("File missing from deploy: %s", file.Name())
		}
	}
}

func TestCopyDir(t *testing.T) {
	//create directory1
	err := os.Mkdir("./directory1", 0777)
	if err != nil {
		t.Error(err)
	}

	//create directory2
	err = os.Mkdir("./directory2", 0777)
	if err != nil {
		t.Error(err)
	}

	//create test file and write a string
	file, err := os.Create("./directory1/testFile.txt")
	if err != nil {
		t.Error(err)
	}
	fmt.Fprint(file, "test")

	//copy directory1 contents to directory2
	copyDir("./directory1", "./directory2")

	//check validity of file in directory2
	copied, err := os.Open("./directory2/testFile.txt")
	if err != nil {
		t.Errorf("Test file was not copied to directory 2, %s", err.Error())
	}
	content, err := ioutil.ReadAll(copied)
	if err != nil {
		t.Errorf("Copied file could not be read, %s", err.Error())
	}
	if string(content) != "test" {
		t.Errorf("Copied file contents do not match original. Expecting 'test', got '%s'", string(content))
	}

	os.RemoveAll("./directory1")
	os.RemoveAll("./directory2")
}

func TestFixFileSystem(t *testing.T) {

	os.MkdirAll("./charts/"+testProfile.Name+"/deploy/templates", 0777)
	fixFileSystem(testProfile)

	//check if helm generated files still exists
	files, err := ioutil.ReadDir("./charts/" + testProfile.Name + "/deploy")
	if err != nil {
		t.Errorf("Error reading test directory: %s", err.Error())
	}
	for _, file := range files {
		if file.IsDir() {
			t.Errorf("Found unexpected directory in chart foler: %s", file.Name())
		}
	}
}
