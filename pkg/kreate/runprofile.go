package kreate

import "fmt"

/*
## kreate run <profile name>
1. Prerequisite: Helm v2.16.3 must be installed and helm init must be ran (confirm)
2. Check: If the istio environment is already deployed to the cluster, there is no need to redeploy it.
    - If istio is not deployed, istio v1.4.5 must be deployed.
3. The custom chart must be created using Part 1. of the the 'Kreate chart' logic.
4. Check: If a custom chart is already deployed to the cluster, it must be upgraded to the new chart.
    - If a custom chart is not already deployed, the new chart must be installed rather than upgraded. 
	- Possible lead for the implementing developer: https://github.com/helm/helm/issues/3353
*/

func RunProfile(name string) error { // current logic was written prior to the 3/3/20 MVP meeting
	fullChartPath := chartsLocation + name
	fmt.Printf("Please edit your Values.yaml file with your favorite text edit @ %s", fullChartPath)
	return nil
}
