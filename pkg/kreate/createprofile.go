package kreate

import (
	"log"

	"helm.sh/helm/v3/pkg/chartutil"
)

// CreateProfile will take in a path to a set up yaml, and then create the output chart in the user's current directory.
func CreateProfile(name string) error {
	// Use Helm code here to create charts for MVP
	fullChartPath, createErr := chartutil.Create(name, chartsLocation)
	if createErr != nil {
		log.Panicf("Error creating chart directory => %+v", createErr)
	}

	// This function merely should copy our templates from /config/templates/
	// the the users /var/local/kreate/<helm-directory-name>/templates/
	// TODO: THIS FUNCTION IS EMPTY AND NEEDS TO BE IMPLEMENTED!!!!
	addRevatureTemplates(fullChartPath)
	// addRevatureValueYaml(fullChartPath)

	return nil
}
