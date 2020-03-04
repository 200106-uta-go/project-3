package kreate

import (
	"log"
	"os"
)

// This value will determine where the helm directories will go by default.
const (
	MOULDFOLDERS = "/var/local/kreate/" // Initial value on where to store helm charts
	PROFILES    = "/etc/kreate/"
)

var (
	chartsLocation string
)

/*
## kreate init
1. Setup kreate's folders to the proper paths (var/local/kreate holds the istio and custom moulds. etc/kreate/ holds profile .yaml files.)
2. Setup kreate's environment variables (If any)
*/

func Init() { // current logic was written prior to the 3/3/20 MVP meeting

	pathErr := os.MkdirAll(MOULDFOLDERS, 1777)
	if pathErr != nil {
		log.Panicf("Error making directory %s => %v", MOULDFOLDERS, pathErr)
	}

	// Register the path of where the helm charts will be stored by checking the env var "KREATE_DATA"
	// Set "KREATE_DATA" == The default value of /var/local/kreate/
	var ok bool
	chartsLocation, ok = os.LookupEnv("KREATE_DATA")
	if !ok {
		setErr := os.Setenv("KREATE_DATA", MOULDFOLDERS)
		if setErr != nil {
			log.Panicf("Error Setting KREATE_DATA to default value => %+v", setErr)
		}
	}

}
