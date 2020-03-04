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
	Name      string   `yaml:"name"`
	ImageURL  string   `yaml:"imageurl"`
	Ports     []string `yaml:"ports"`
	Endpoints []string `yaml:"endpoints"`
}

profilePath = "/etc/kreate/" 

//GetProfile gets the profile file and return the data as a struct
func GetProfile(profileName string) Profile {
	//open profile
	file, err := os.Open(profilePath + profileName)
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
