package kreate

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
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

// This is already defined by PROFILES const in initialization.go. This is where the vars and consts are defined.

//GetProfile gets the profile file and return the data as a struct
func GetProfile(profileName string) Profile {
	//open profile
	file, err := os.Open(PROFILES + profileName)
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
