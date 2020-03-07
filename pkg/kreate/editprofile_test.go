package kreate

import (
	"log"
	"os"
	"testing"
)

// This value will determine where the helm directories will go by default.
const (
	TESTINGPROFILEDIRECTORY = "/etc/kreate/"
)

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
			Ports:       []string{"80", "8080"},
			Endpoints:   []string{"/", "/helloworldSecondApp"},
		},
	},
}

func init() {
	pathErr := os.MkdirAll(TESTINGPROFILEDIRECTORY, 1777)
	if pathErr != nil {
		log.Panicf("Error making directory %s => %v", PROFILES, pathErr)
	}
}

// CreateTestingProfile will take a name defined by the user and then ouput a default file with the users
// default editor.
func CreateProfile(name string) error {
	// Check if given profile name exists
	if _, err := os.Stat(PROFILES + name + ".yaml"); err != nil {
		// If profile is not exist, create new yaml file
		file, err := os.Create(PROFILES + name + ".yaml")
		if err != nil {
			return err
		}
		defer file.Close()

		// Marshal defaultProfile struct
		defaultProfile.Name = name
		bytes, err := yaml.Marshal(defaultProfile)
		if err != nil {
			return err
		}
		_, err = file.Write(bytes)
		if err != nil {
			return err
		}
		defaultProfile.Name = "myProfileName"
		// Open generated yaml file with text editor
	} else {
		return err
	}
	return nil


func TestEditProfile(t *testing.T) {
}

func ExampleEditProfile() {

}
