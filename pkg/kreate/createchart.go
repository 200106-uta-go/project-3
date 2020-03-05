package kreate

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

/*
## kreate chart <profile name>
1. Part 1. (to be reused within kreate run)
    - The specified profile must be converted from its profile format to the helm chart's values.yaml format
    - The created values.yaml must then be copied to the kustom chart.
2. Part 2.
    - The kustom chart must then be copied to a user-friendly directory (implementation as descretioned by developer)
*/

//CreateChart creates a helm chart using the data provided in profile
func CreateChart(profileName string) {
	profile := GetProfile(profileName + ".yaml")

	createValues(profile)

	tmpl, err := os.Open("chart.tmpl")
	if err != nil {
		panic(err)
	}

	//chart template and values.yaml should be created at this point

	//add values into chart for deployment yaml
	populateChart("values.yaml", tmpl)
}

//createValues creates a values.yaml based on a profile
func createValues(profile Profile) {
	//create values yaml
	file, err := os.Create("values.yaml")
	if err != nil {
		panic(err)
	}

	bytes, err := yaml.Marshal(profile)
	if err != nil {
		panic(err)
	}

	// Edited by CreateProfile Team so our code can run. Feel free to undo this.
	// 	written, err := file.Write(bytes)	if written == 0 {
	// 		panic("Nothing was written to values.yaml")
	// 	}
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}

//populateChart injects the values inside filename into a chart template
func populateChart(filename string, template *os.File) {
	fmt.Println("Test Complete")
}
