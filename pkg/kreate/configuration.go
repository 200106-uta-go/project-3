package kreate

import (
	"log"
	"os"
)

// This value will determine where the helm directories will go by default.
const (
	HELMFOLDERS = "/var/local/kreate" // Initial value on where to store helm charts
)

var (
	chartPath = ""
)

// Init will set up the environment variables for use with Kreate
func Init() {

	// Register the path of where the helm charts will be stored by checking the env var "KREATE_DATA"
	// Set "KREATE_DATA" == The default value of /var/local/kreate
	var ok bool
	chartPath, ok = os.LookupEnv("KREATE_DATA")
	if !ok {
		setErr := os.Setenv("KREATE_DATA", HELMFOLDERS)
		if setErr != nil {
			log.Panicf("Error Setting KREATE_DATA to default value => %+v", setErr)
		}
	}

}
