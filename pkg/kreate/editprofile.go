package kreate

import (
	"flag"
	"fmt"
	"log"
)

/*
## kreate edit <profile name>
1. The selected profile must be opened for modification.
2. After modification, the profile will save.
*/

//Global var, which are our flags
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
	flag.StringVar(&Profile_ClusterPort, "pcp", "", "ClusterIp for config.")

	flag.StringVar(&Profile_App_Name, "pan", "", "Under App, the Name value.")
	flag.StringVar(&Profile_App_ImageURL, "pai", "", "Under App, the ImageURL.")
	flag.StringVar(&Profile_App_ServiceName, "pas", "", "Under App, the ServiceName value.")
	flag.IntVar(&Profile_App_ServicePort, "pasp", 0, "Under App, the ServicePort value.")
	flag.StringVar(&Profile_App_Port, "pap", "", "Under App, Port Value.")
	flag.StringVar(&Profile_App_Endpoint, "pae", "", "Under App, Endpoint Value.")

	flag.Parse()
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

// EditProfile is a function that checks a single cluster and overwrites any profile information, while checking through app specific information and adjusting according to provided flags
func EditProfile(pf Profile) (Profile, error) { // current logic was written prior to the 3/3/20 MVP meeting
	//fullChartPath := chartsLocation + name
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
	noImageURL := false
	noServiceName := false
	noServicePort := false
	noPort := false
	noEndpoint := false

	CheckAppValues(&noImageURL, &noServiceName, &noServicePort, &noPort, &noEndpoint)

	fmt.Println(Profile_App_Name)

	if Profile_App_Name == "" {
		//ALL GOOD, no app is being changed.
		if noImageURL && noServiceName && noServicePort && noPort && noEndpoint {
			return pf, nil
		}
		// Editing app without specificing app.Name, program can not determine which
		//	app to change so values are unchanged and log an error to the user
		log.Print("Editing app without specificing app.Name, program can not determine which app to change so values are unchanged.")
		return pf, nil

	} else if Profile_App_Name != "" {
		fmt.Println("elseIF")
		fmt.Println(pf.Apps)
		for i, _ := range pf.Apps {
			// Profile_App_Name matches with an existing app_name
			fmt.Println("pf.Apps[i].Name, Profile_App_Name")
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
			return pf, nil
		}
		// Profile_App_Name does not match with an existing app_name
		log.Print("Editing app an App_Name that does not exist, program can not modify an app that does not exist.")
		return pf, nil
	}
	return pf, nil
}
