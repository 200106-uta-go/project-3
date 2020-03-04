package kreate

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	yaml "gopkg.in/yaml.v3"
)

func TestCreateProfile(t *testing.T) {
	var myName = "gopher"

	CreateProfile(myName)

	expected := Profile{
		Name:         "gopher",
		ClusterName:  "Your Cluster Name",
		ClusterIP:    "127.0.0.1",
		ClusterPorts: []string{"80"},
		Apps: []App{
			App{
				Name:      "helloWorld",
				ImageURL:  "https://hub.docker.com/hello-world",
				Ports:     []string{"80", "8080"},
				Endpoints: []string{"/", "/helloworld"},
			},
		},
	}

	var actual Profile

	data, _ := ioutil.ReadFile(myName)

	yaml.Unmarshal(data, &actual)

	if reflect.DeepEqual(actual, expected) {

		t.Error(fmt.Sprintf("Actual did not match expected, Actual => %+v \nExpected => %+v\n", actual, expected))
	}

	fmt.Println(reflect.DeepEqual(actual, expected))
	// Output : true
}
