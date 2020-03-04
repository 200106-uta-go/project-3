package kreate

import (
	"log"
	"os"
)

// This value will determine where the helm directories will go by default.
const (
	HELMFOLDERS = "/var/local/kreate/" // Initial value on where to store helm charts
	PROFILES    = "/etc/kreate/"
)

var (
	chartsLocation string
)

// Init will set up the environment variables for use with Kreate
func Init() {

	pathErr := os.MkdirAll(HELMFOLDERS, 1777)
	if pathErr != nil {
		log.Panicf("Error making directory %s => %v", HELMFOLDERS, pathErr)
	}

	// Register the path of where the helm charts will be stored by checking the env var "KREATE_DATA"
	// Set "KREATE_DATA" == The default value of /var/local/kreate/
	var ok bool
	chartsLocation, ok = os.LookupEnv("KREATE_DATA")
	if !ok {
		setErr := os.Setenv("KREATE_DATA", HELMFOLDERS)
		if setErr != nil {
			log.Panicf("Error Setting KREATE_DATA to default value => %+v", setErr)
		}
	}

}
