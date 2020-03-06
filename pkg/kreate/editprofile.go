package main

import (
	"fmt"

	kre8 "github.com/200106-uta-go/project-3/pkg/kreate"
)

func main() {
	kre8.Initialization()
	// Call GetProfile to obtain a profile struct from a given file name defined by the -f flag.
	file, _ := kre8.CreateProfile("test123")
	// Call out EditProfile function.
	profileStruct := kre8.GetProfile("test123.yaml")

	fmt.Println(profileStruct)
	//f, _ := os.OpenFile("/etc/kreate/test123.yaml", os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()

	profileStruct, _ = kre8.EditProfile(profileStruct, "test123.yaml")

	fmt.Println(profileStruct)

}

----------------

package kreate

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

/*
## kreate edit <profile name>
1. The selected profile must be opened for modification.
2. After modification, the profile will save.
*/

// YamlFileName is th name of the yaml.
var YamlFileName string = "CustomProfile.yaml"

// Flag variables that can be configured.

var FileName string

var Profile_Name string
var Profile_ClusterName string
var Profile_ClusterIP string
var Profile_ClusterPort string

var Profile_App_Name string
var Profile_App_ImageURL string
var Profile_App_ServiceName string
var Profile_App_ServicePort int
var Profile_App_Port string
var Profile_App_Endpoint string

func init() {
	flag.StringVar(&FileName, "f", "testProfile.yaml", "Default is named testProfile.yaml, this is located within /etc/kreate.")

	flag.StringVar(&Profile_Name, "pn", "", "Name for config.")
	flag.StringVar(&Profile_ClusterName, "pcn", "", "ClusterName for config.")
	flag.StringVar(&Profile_ClusterIP, "pci", "", "ClusterIp for config.")
	flag.StringVar(&Profile_ClusterPort, "pcp", "", "ClusterPort for config.")

	flag.StringVar(&Profile_App_Name, "pan", "", "Under App, the Name value.")
	flag.StringVar(&Profile_App_ImageURL, "pai", "", "Under App, the ImageURL.")
	flag.StringVar(&Profile_App_ServiceName, "pas", "", "Under App, the ServiceName value.")
	flag.IntVar(&Profile_App_ServicePort, "pasp", 0, "Under App, the ServicePort value.")
	flag.StringVar(&Profile_App_Port, "pap", "", "Under App, Port Value.")
	flag.StringVar(&Profile_App_Endpoint, "pae", "", "Under App, Endpoint Value.")

	flag.Parse()
}

// Check panics on any error that is not nill.
func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// ProfileToYaml ...
func ProfileToYaml(pf Profile) error {
	Check(os.Remove(PROFILES + YamlFileName))
	f, err := os.OpenFile(PROFILES+YamlFileName, os.O_RDWR|os.O_CREATE, 0755)
	Check(err)

	bytes, err := yaml.Marshal(&pf)
	if err != nil {
		return err
	}
	fmt.Println("Finish marshal")

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	fmt.Println("Finish write")
	return nil
}

// CheckAppValues will determine if any App information is passed into the flag and change its boolean value to determine future logic flow.
func CheckAppValues(noImageURL, noServiceName, noServicePort, noPort, noEndpoint *bool) {
	if Profile_App_ImageURL == "" {
		*noImageURL = true
	}
	if Profile_App_ServiceName == "" {
		*noServiceName = true
	}
	if Profile_App_ServicePort == 0 {
		*noServicePort = true
	}
	if Profile_App_Port == "" {
		*noPort = true
	}
	if Profile_App_Endpoint == "" {
		*noEndpoint = true
	}
}

// EditProfile is a function that Checks a single cluster and overwrites any profile information, while Checking through app specific information and adjusting according to provided flags
func EditProfile(pf Profile, YamlName string) (Profile, error) { // current logic was written prior to the 3/3/20 MVP meeting
	YamlFileName = YamlName
	noImageURL := false
	noServiceName := false
	noServicePort := false
	noPort := false
	noEndpoint := false

	if Profile_Name != "" {
		pf.Name = Profile_Name
	}
	if Profile_ClusterName != "" {
		pf.ClusterName = Profile_ClusterName
	}
	if Profile_ClusterIP != "" {
		pf.ClusterIP = Profile_ClusterIP
	}
	if Profile_ClusterPort != "" {
		foundPort := false
		for index := range pf.ClusterPorts {
			if pf.ClusterPorts[index] == Profile_ClusterPort {
				foundPort = true
			}
		}
		if foundPort == false {
			pf.ClusterPorts = append(pf.ClusterPorts, Profile_ClusterPort)
		}
	}
	CheckAppValues(&noImageURL, &noServiceName, &noServicePort, &noPort, &noEndpoint)
	if Profile_App_Name == "" {
		//ALL GOOD, no app is being changed.
		if noImageURL && noServiceName && noServicePort && noPort && noEndpoint {
			Check(ProfileToYaml(pf))
			return pf, nil
		}
		// Editing app without specificing app.Name, program can not determine which
		//	app to change so values are unchanged and log an error to the user
		log.Print("Editing app without specificing app.Name, program can not determine which app to change so values are unchanged.")
		Check(ProfileToYaml(pf))
		return pf, nil
	} else if Profile_App_Name != "" {
		for i, _ := range pf.Apps {
			// Profile_App_Name matches with an existing app_name
			if pf.Apps[i].Name == Profile_App_Name {
				if noImageURL == false {
					pf.Apps[i].ImageURL = Profile_App_ImageURL
				}
				if noServiceName == false {
					pf.Apps[i].ServiceName = Profile_App_ServiceName
				}
				if noServicePort == false {
					pf.Apps[i].ServicePort = Profile_App_ServicePort
				}
				if noPort == false {
					pf.Apps[i].Ports = append(pf.Apps[i].Ports, Profile_App_Port)
				}
				if noEndpoint == false {
					pf.Apps[i].Endpoints = append(pf.Apps[i].Endpoints, Profile_App_Endpoint)
				}
			}
			// ALL GOOD, no app is being changed
			Check(ProfileToYaml(pf))
			return pf, nil
		}
		// Profile_App_Name does not match with an existing app_name
		log.Print("Editing app an App_Name that does not exist, program can not modify an app that does not exist.")
		Check(ProfileToYaml(pf))
		return pf, nil
	}
	Check(ProfileToYaml(pf))
	return pf, nil
}
