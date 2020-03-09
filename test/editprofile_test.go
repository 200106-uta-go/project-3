package kreate

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/200106-uta-go/project-3/pkg/kreate"
	"gopkg.in/yaml.v2"
)

// This value will determine where the helm directories will go by default.
const (
	TESTINGPROFILEDIRECTORY = "/etc/kreate_test/"
)

// Profile defines the profile struct and converts between the struct format and the yaml format
type Profile struct {
	Name         string   `yaml:"name"`
	ClusterName  string   `yaml:"clustername"`
	ClusterIP    string   `yaml:"clusterip"`
	ClusterPorts []string `yaml:"clusterports"`
	Apps         []App    `yaml:"apps"`
}

// App ...
type App struct {
	Name        string   `yaml:"name"`
	ImageURL    string   `yaml:"imageurl"`
	ServiceName string   `yaml:"servicename"`
	ServicePort int      `yaml:"serviceport"`
	Ports       []string `yaml:"ports"`
	Endpoints   []string `yaml:"endpoints"`
}

var defaultTestingProfile *Profile = &Profile{
	Name:         "myProfileName",
	ClusterName:  "Your Cluster Name",
	ClusterIP:    "127.0.0.1",
	ClusterPorts: []string{"80"},
	Apps: []App{
		App{
			Name:        "helloWorld",
			ImageURL:    "https://hub.docker.com/hello-world",
			ServiceName: "hello-service",
			ServicePort: 7777,
			Ports:       []string{"80", "8080"},
			Endpoints:   []string{"/", "/helloworld"},
		},
		App{
			Name:        "helloWorldSecondApp",
			ImageURL:    "https://hub.docker.com/hello-world",
			ServiceName: "hello-service-Second",
			ServicePort: 7778,
			Ports:       []string{"90", "9090"},
			Endpoints:   []string{"/", "/helloworldSecondApp"},
		},
	},
}

func init() {
	pathErr := os.MkdirAll(TESTINGPROFILEDIRECTORY, 1777)
	if pathErr != nil {
		log.Panicf("Error making directory %s => %v", TESTINGPROFILEDIRECTORY, pathErr)
	}
}

// check function determines if the input error is a non nill value and performs a panic if so.
func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// CreateTestingProfile will take a name defined by the user and then ouput a default file with the users
// default editor.
func CreateTestingProfile(name string) error {
	// Check if given profile name exists
	if _, err := os.Stat(TESTINGPROFILEDIRECTORY + name + ".yaml"); err != nil {
		// If profile is not exist, create new yaml file
		file, err := os.Create(TESTINGPROFILEDIRECTORY + name + ".yaml")
		if err != nil {
			return err
		}
		defer file.Close()

		// Marshal defaultTestingProfile struct
		defaultTestingProfile.Name = name
		bytes, err := yaml.Marshal(defaultTestingProfile)
		if err != nil {
			return err
		}
		_, err = file.Write(bytes)
		if err != nil {
			return err
		}
		defaultTestingProfile.Name = "myProfileName"
		// Open generated yaml file with text editor
	} else {
		return err
	}
	return nil
}

//GetTestProfile gets the profile file and return the data as a struct
func GetTestProfile(profileName string) Profile {
	//check if profileName has an extension, if not add .yaml
	if !strings.HasSuffix(profileName, ".yaml") && !strings.HasSuffix(profileName, ".yml") {
		profileName += ".yaml"
	}

	//open profile
	file, err := os.Open(TESTINGPROFILEDIRECTORY + profileName)
	if err != nil {
		panic(err)
	}

	//read all data in profile
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	//unmarshal file's contents into profile struct
	profile := Profile{}
	yaml.Unmarshal(bytes, &profile)

	return profile
}

func TestEditProfile(t *testing.T) {
	check(CreateTestingProfile("defaultTest"))
	//pf := GetTestProfile("defaultTest")
	pf, err := kreate.EditProfile("defaultTest")
	check(err)
	if pf.Name != "NEWNAME" {
		t.Error("Profile struct and yaml did not change to NEWNAME.")
	}

}

func ExampleEditProfile() {

}
