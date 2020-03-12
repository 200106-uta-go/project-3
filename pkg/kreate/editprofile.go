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
	"strings"

	"gopkg.in/yaml.v2"
)

// YamlFileName is a global variable that holds the title of the yaml file
// 	that will be used to create the yaml file.
var YamlFileName string

//FLAG VARIABLES
var Name string
var ClusterName string
var ClusterIP string
var ClusterPort string

var AppName string
var AppImageURL string
var AppServiceName string
var AppServicePort int
var AppPort string
var AppEndpoint string

// init assigns flag defaults and parses the flags, which are used to assign new values if they are set.
// JZ: moved flag parsing to EditProfile since an init func here means an init func for everything within the kreate package
func init() {

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

	// Write the new updated structure to the file specified by YamlFileName.
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

// CheckAppValues will determine if any App information is set to non empty value and
// 	change a boolean value, which is used to determine whether to assign values or log a report depending on further logic.
func CheckAppValues(noImageURL, noServiceName, noServicePort, noPort, noEndpoint *bool) {
	// This 5 if statements are checking corresponding flags to know what varaibles to change in the file pointed to by corresponding.
	if AppImageURL == "" {
		*noImageURL = true
	}
	if AppServiceName == "" {
		*noServiceName = true
	}
	if AppServicePort == 0 {
		*noServicePort = true
	}
	if AppPort == "" {
		*noPort = true
	}
	if AppEndpoint == "" {
		*noEndpoint = true
	}
}

// EditProfile is a function that Checks a single cluster and overwrites any profile information, while Checking through app specific information
//	 and adjusting according to provided flags
func EditProfile(YamlName string) (Profile, error) {
	// Shift arguments back 2 to parse flags from os.Args[1:]
	os.Args = os.Args[2:]
	
	flag.StringVar(&Name, "name", "", "Sets the name of profile.")
	flag.StringVar(&ClusterName, "clustername", "", "Sets the clustername of the profile.")
	flag.StringVar(&ClusterIP, "clusterip", "", "Sets the clusterip of the profile.")
	flag.StringVar(&ClusterPort, "clusterport", "", "Append a clusterport to the profile.")

	flag.StringVar(&AppName, "NameOfApp", "", "Specifies the name of the app which will be modified by the app-related input flags.")
	flag.StringVar(&AppImageURL, "imageurl", "", "An App-related flag. Sets the imageurl of the App specified by the NameOfApp flag.")
	flag.StringVar(&AppServiceName, "servicename", "", "An App-related flag. Sets the servicename of the App specified by the NameOfApp flag.")
	flag.IntVar(&AppServicePort, "serviceport", 0, "An App-related flag. Sets the serviceport of the App specified by the NameOfApp flag.")
	flag.StringVar(&AppPort, "port", "", "An App-related flag. Appends a port to the App specified by the NameOfApp flag.")
	flag.StringVar(&AppEndpoint, "endpoint", "", "An App-related flag. Appends an endpoint to the App specified by the NameOfApp flag.")

	flag.Parse()
	if (YamlName == "") {
		fmt.Println("No profile specified")
		os.Exit(1)
	} else if (!strings.HasSuffix(YamlName, ".yaml") && !strings.HasSuffix(YamlName, ".yml")) {
		YamlName += ".yaml"
	}

	YamlFileName = YamlName
	pf := GetProfile(YamlFileName)
	noImageURL := false
	noServiceName := false
	noServicePort := false
	noPort := false
	noEndpoint := false

	if flag.NFlag() == 0 {
	 	OpenFileInEditor(YamlName)
		return pf, nil 
	}

	//Checking this 4 if statements to see if any one of our next four flags where set to make changes to the file pointed to by YamlFileName.
	if Name != "" {
		pf.Name = Name
	}
	if ClusterName != "" {
		pf.ClusterName = ClusterName
	}
	if ClusterIP != "" {
		pf.ClusterIP = ClusterIP
	}
	// This is an array of ports so we are chekcing if the flag to add aport has been declared and specified.
	if ClusterPort != "" {
		foundPort := false
		for index := range pf.ClusterPorts {
			if pf.ClusterPorts[index] == ClusterPort {
				foundPort = true
			}
		}
		if foundPort == false {
			pf.ClusterPorts = append(pf.ClusterPorts, ClusterPort)
		}
	}
	//Here used boolean and to see what flags where specified to edit the yaml structure file specified by YamlFileName.
	CheckAppValues(&noImageURL, &noServiceName, &noServicePort, &noPort, &noEndpoint)
	if AppName == "" {
		//ALL GOOD, no app is being changed.
		if noImageURL && noServiceName && noServicePort && noPort && noEndpoint {
			Check(ProfileToYaml(pf))
			str, _ := shellCommand(fmt.Sprintf("cat %s", YamlName), PROFILES)
			fmt.Println(str)
			return pf, nil
		}
		// Editing app without specificing app.Name, program can not determine which
		//	app to change so values are unchanged and log an error to the user
		log.Print("Editing app without specificing app.Name, program can not determine which app to change so values are unchanged.")
		Check(ProfileToYaml(pf))
		str, _ := shellCommand(fmt.Sprintf("cat %s", YamlName), PROFILES)
		fmt.Println(str)
		return pf, nil
	} else if AppName != "" {
		for i := 0; i < len(pf.Apps); i++ {
			// AppName is an array of appname so we must cross reference to see if a specific app name
			//	is mentions in order to modify desired variables of that specific appname.
			if pf.Apps[i].Name == AppName {
				if noImageURL == false {
					pf.Apps[i].ImageURL = AppImageURL
				}
				if noServiceName == false {
					pf.Apps[i].ServiceName = AppServiceName
				}
				if noServicePort == false {
					pf.Apps[i].ServicePort = AppServicePort
				}
				if noPort == false {
					pf.Apps[i].Ports = append(pf.Apps[i].Ports, AppPort)
				}
				if noEndpoint == false {
					pf.Apps[i].Endpoints = append(pf.Apps[i].Endpoints, AppEndpoint)
				}
			}
			//Case: no app name, and no app details proved, just return any changes that happen on clsuter details.
			Check(ProfileToYaml(pf))
			str, _ := shellCommand(fmt.Sprintf("cat %s", YamlName), PROFILES)
			fmt.Println(str)
			return pf, nil
		}
		// AppName does not match with an existing App.Names.
		log.Print("Editing app an App_Name that does not exist, program can not modify an app that does not exist.")
		Check(ProfileToYaml(pf))
		str, _ := shellCommand(fmt.Sprintf("cat %s", YamlName), PROFILES)
		fmt.Println(str)
		return pf, nil
	}

	Check(ProfileToYaml(pf))
	str, _ := shellCommand(fmt.Sprintf("cat %s", YamlName), PROFILES)
	fmt.Println(str)
	return pf, nil
}
