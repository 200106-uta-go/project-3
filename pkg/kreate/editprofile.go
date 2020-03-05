package kreate

import "fmt"

/*
## kreate edit <profile name>
1. The selected profile must be opened for modification.
2. After modification, the profile will save.
*/

func EditProfile(name string) error {  // current logic was written prior to the 3/3/20 MVP meeting
	fullChartPath := chartsLocation + name
	fmt.Printf("Please edit your Values.yaml file with your favorite text edit @ %s", fullChartPath)
	return nil
}
