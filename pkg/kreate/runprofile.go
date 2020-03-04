package kreate

import "fmt"

// RunProfile will allow the user to edit the values Yaml created by helm.
// How we choose to do this is up for discussion.
func RunProfile(name string) error {
	fullChartPath := chartsLocation + name
	fmt.Printf("Please edit your Values.yaml file with your favorite text edit @ %s", fullChartPath)
	return nil
}
