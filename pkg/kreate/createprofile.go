package kreate

import (
	"os"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

/*
## kreate profile <profile name>
1. The following variables must be saved to the profile .yaml file.
    - A profile name.
    - The foreign cluster's name (the name of the fallback cluster)
    - The foreign cluster's IP
    - The foreign cluster's Port of Entry
    - The business application's (the client's application container) image URL
    - An array of the business application's container exposable ports (non-specific)
    - An array of the business application's endpoints (/players, /info, /static. /, ect.)
2. Created profiles will be saved to etc/kreate/
*/

var defaultProfile *Profile = &Profile{
	Name:         "myprofile",
	ClusterName:  "foreign-cluster",
	ClusterIP:    "127.0.0.1:30101",
	ClusterPorts: []string{"30101"},
	Apps: []App{
		App{
			Name:        "hello-world",
			ImageURL:    "hello-world",
			ServiceName: "hello-service",
			ServicePort: 7777,
			Ports:       []string{"80", "8080"},
			Endpoints:   []string{"/", "/helloworld"},
		},
	},
}

// DefaultEditor is vim because we're adults ;)
const DefaultEditor = "nano"

// CreateProfile will take a name defined by the user and then ouput a default file with the users
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
		defaultProfile.Name = "myprofile"
		// Open generated yaml file with text editor
	} else {
		return err
	}
	return nil
}

// OpenFileInEditor opens filename in a text editor.
func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(filename, ".yaml") && !strings.HasSuffix(filename, ".yml") {
		filename += ".yaml"
	}

	cmd := exec.Command(executable, PROFILES+filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
