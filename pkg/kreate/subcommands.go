package kreate

import (
	"fmt"
	"log"

	"helm.sh/helm/v3/pkg/chartutil" // be sure to run 'go get -d -u -v helm.sh/helm/v3/...' in order to grab all of helm's dependencies
)

// import the helm code here to implement the function

// CreateChart will take in a path to a set up yaml, and then create the output chart in the user's current directory.
func CreateChart(name string) error {
	// Use Helm code here to create charts for MVP
	fullChartPath, createErr := chartutil.Create(name, chartsLocation)
	if createErr != nil {
		log.Panicf("Error creating chart directory => %+v", createErr)
	}

	addRevatureTemplates(fullChartPath)
	// addRevatureValueYaml(fullChartPath)

	return nil
}

// EditValues will allow the user to edit the values Yaml created by helm.
// How we choose to do this is up for discussion.
func EditValues(name string) error {
	fullChartPath := chartsLocation + name
	fmt.Printf("Please edit your Values.yaml file with your favorite text edit @ %s", fullChartPath)
	return nil
}
