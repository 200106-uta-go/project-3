package kreate

import (
	"log"
	"os"
	"os/exec"

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

var defaultProfile Profile = Profile{
	Name:         "myProfileName",
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

// DefaultEditor is vim because we're adults ;)
const DefaultEditor = "nano"

// CreateProfile will take a name defined by the user and then ouput a default file with the users
// default editor.
func CreateProfile(name string) {
	// Check if given profile name exists
	if _, err := os.Stat(name + ".yaml"); err != nil {
		// If profile is not exist, create new yaml file
		file, err := os.Create(name + ".yaml")
		if err != nil {
			log.Panicln(err)
		}
		defer file.Close()

		// Marshal defaultProfile struct
		bytes, err := yaml.Marshal(defaultProfile)
		_, err = file.Write(bytes)
	}

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

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
