package kreate

/*
Authors: Joshua Nguyen, and Hector Moreno.
Date: March 04, 2020.
Section: UTA - Go batch.
Trainer: Mehrab R.
*/

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//FLAG VARIABLES
var profileName string
var profileClusterName string
var profileClusterIP string
var profileClusterPort string

var profileAppName string
var profileAppImageURL string
var profileAppServiceName string
var profileAppServicePort int
var profileAppPort string
var profileAppEndpoint string

// init assigns flag defaults and parses the flags, which are used to assign new values if they are set.
func init() {

	flag.StringVar(&profileName, "profileName", "", "Name for config.")
	flag.StringVar(&profileClusterName, "profileClusterName", "", "ClusterName for config.")
	flag.StringVar(&profileClusterIP, "profileClusterIP", "", "ClusterIp for config.")
	flag.StringVar(&profileClusterPort, "profileClusterPort", "", "ClusterPort for config.")

	flag.StringVar(&profileAppName, "profileAppName", "", "Under App, the Name value.")
	flag.StringVar(&profileAppImageURL, "profileAppImageURL", "", "Under App, the ImageURL.")
	flag.StringVar(&profileAppServiceName, "profileAppServiceName", "", "Under App, the ServiceName value.")
	flag.IntVar(&profileAppServicePort, "profileAppServicePort", 0, "Under App, the ServicePort value.")
	flag.StringVar(&profileAppPort, "profileAppPort", "", "Under App, Port Value.")
	flag.StringVar(&profileAppEndpoint, "profileAppEndpoint", "", "Under App, Endpoint Value.")

	flag.Parse()
}

// Check is a function that panics on any error that is not nill. Used to condense error handling into a function call.
func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// ProfileToYaml takes in a profile struct and delete old yaml file with the same name as the 
// 	value set in YamlFileName, and create a new yaml file with that same YamlFileName. 
func ProfileToYaml(pf Profile) error {
	Check(os.Remove(PROFILES + YamlFileName))
	f, err := os.OpenFile(PROFILES+YamlFileName, os.O_RDWR|os.O_CREATE, 0755)
	Check(err)
	//Turns the structure into yaml format.
	bytes, err := yaml.Marshal(&pf)
	if err != nil {
		return err
	}
	fmt.Println("Finish marshal")
	// Write the new updated structure to the file specified by YamlFileName. 
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	fmt.Println("Finish write")
	return nil
}

// CheckAppValues will determine if any App information is set to non empty value and 
// 	change a boolean value, which is used to determine whether to assign values or log a report depending on further logic.
func CheckAppValues(noImageURL, noServiceName, noServicePort, noPort, noEndpoint *bool) {
	// This 5 if statements are checking corresponding flags to know what varaibles to change in the file pointed to by corresponding.
	if profileAppImageURL == "" {
		*noImageURL = true
	}
	if profileAppServiceName == "" {
		*noServiceName = true
	}
	if profileAppServicePort == 0 {
		*noServicePort = true
	}
	if profileAppPort == "" {
		*noPort = true
	}
	if profileAppEndpoint == "" {
		*noEndpoint = true
	}
}

// EditProfile is a function that Checks a single cluster and overwrites any profile information, while Checking through app specific information
//	 and adjusting according to provided flags
func EditProfile(pf Profile, YamlName string) (Profile, error) { // current logic was written prior to the 3/3/20 MVP meeting
	YamlFileName = YamlName
	noImageURL := false
	noServiceName := false
	noServicePort := false
	noPort := false
	noEndpoint := false
	//Checking this 4 if statements to see if any one of our next four flags where set to make changes to the file pointed to by YamlFileName.
	if profileName != "" {
		pf.Name = profileName
	}
	if profileClusterName != "" {
		pf.ClusterName = profileClusterName
	}
	if profileClusterIP != "" {
		pf.ClusterIP = profileClusterIP
	}
	// This is an array of ports so we are chekcing if the flag to add aport has been declared and specified.
	if profileClusterPort != "" {
		foundPort := false
		for index := range pf.ClusterPorts {
			if pf.ClusterPorts[index] == profileClusterPort {
				foundPort = true
			}
		}
		if foundPort == false {
			pf.ClusterPorts = append(pf.ClusterPorts, profileClusterPort)
		}
	}
	//Here used boolean and to see what flags where specified to edit the yaml structure file specified by YamlFileName.
	CheckAppValues(&noImageURL, &noServiceName, &noServicePort, &noPort, &noEndpoint)
	if profileAppName == "" {
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
	} else if profileAppName != "" {
		for i int := 0; i < len(pf.Apps); i++ {
			// profileAppName is an array of appname so we must cross reference to see if a specific app name 
			//	is mentions in order to modify desired variables of that specific appname.
			if pf.Apps[i].Name == profileAppName {
				if noImageURL == false {
					pf.Apps[i].ImageURL = profileAppImageURL
				}
				if noServiceName == false {
					pf.Apps[i].ServiceName = profileAppServiceName
				}
				if noServicePort == false {
					pf.Apps[i].ServicePort = profileAppServicePort
				}
				if noPort == false {
					pf.Apps[i].Ports = append(pf.Apps[i].Ports, profileAppPort)
				}
				if noEndpoint == false {
					pf.Apps[i].Endpoints = append(pf.Apps[i].Endpoints, profileAppEndpoint)
				}
			}
			//Case: no app name, and no app details proved, just return any changes that happen on clsuter details.
			Check(ProfileToYaml(pf))
			return pf, nil
		}
		// profileAppName does not match with an existing App.Names.
		log.Print("Editing app an App_Name that does not exist, program can not modify an app that does not exist.")
		Check(ProfileToYaml(pf))
		return pf, nil
	}
	
	Check(ProfileToYaml(pf))
	return pf, nil
}
