package kreate

import (
	"log"
	"os/exec"
)

// import the helm code here to implement the function

// CreateChart will take in a path to a set up yaml, and then create the output chart in the user's current directory.
func CreateChart(name string) error {
	// Use Helm code here to create charts for MVP

	// Make chart folder in chart path. First make path.
	newChartPath := chartPath + name
	// Change to the /var/local/kreate folder so that all the helm create commands will form a folder here.
	createErr := exec.Command("helm", "create", newChartPath).Run()
	if createErr != nil {
		log.Panicf("Error creating chart directory => %+v", createErr)
	}

	return nil
}
