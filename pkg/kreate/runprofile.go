package kreate

import (
	"fmt"
	"os/exec"
	"strings"
)

/*
## kreate run <profile name>
1. Prerequisite: Helm v2.16.3 must be installed and helm init must be ran (confirm)
2. Check: If the istio environment is already deployed to the cluster, there is no need to redeploy it.
    - If istio is not deployed, istio v1.4.5 must be deployed.
3. The custom chart must be created using Part 1. of the the 'Kreate chart' logic.
4. Check: If a custom chart is already deployed to the cluster, it must be upgraded to the new chart.
    - If a custom chart is not already deployed, the new chart must be installed rather than upgraded.
*/

// RunProfile installs the specified profile and the istio environment in tandem, into the kubernetes cluster.
// If an istio environment is already deployed, RunProfile will not attempt to redeploy it.
// If a profile is already deployed, RunProfile will upgrade the deployment using the new profile.
func RunProfile(profileName string) string {
	currentDir := "./"

	// 1. confirm prerequisites. TODO confirm kreate init is ran prior to allowing kreate run
	// EX: if (kreate.isInitialized()) -> RunProfile()
	// the logic below should work for now
	str, err := shellCommand("sudo helm version", currentDir)
	if !strings.Contains(str, "v2.16.3") && err == nil {
		return fmt.Sprintf("Helm version is not v2.16.3 (required). Run kreate init to install Helm v2.16.3.")
	} else if err != nil {
		// Prerequisite: Confirm Helm init is already ran.
		if strings.Contains(str, "could not find tiller") {
			return fmt.Sprint("Could not find tiller. Run kreate init to initialize helm.")
		}
		// Misc. error (Helm not installed, no Cluster, ect.)
		return fmt.Sprintf("Error: %s", str)
	}
	// Prerequisite: Check if istio is already deployed
	str, err = shellCommand("kubectl get services -n istio-system", currentDir)
	if strings.Contains(str, "No resources found in istio-system namespace") {
		fmt.Println("Istio is not yet deployed. Run kreate init to deploy Istio to the cluster.")
	} else if err != nil {
		// Misc. error (Helm not installed, no Cluster, ect.)
		return fmt.Sprintf("Error: %s", str)
	}

	// 2. Create charts using profile
	CreateChart(profileName)
	profile := GetProfile(profileName)
	releaseName := strings.ToLower(profileName)
	releaseName = strings.ReplaceAll(releaseName, " ", "-")
	
	// 4. Deploy/Upgrade custom chart
	str, err = shellCommand(fmt.Sprintf("helm install -n %s ./charts/%s", releaseName, profile.Name), currentDir)
	if err != nil {
		if strings.Contains(str, fmt.Sprintf("Error: a release named %s already exists.", releaseName)) {
			str, err = shellCommand(fmt.Sprintf("helm upgrade %s ./charts/%s", releaseName, profile.Name), currentDir)
			if err != nil {
				return fmt.Sprintf("Error: Failed to upgrade custom helm chart - %s", str)
			}
			return fmt.Sprintf("Profile %s upgraded successfully", profileName)
		}
		return fmt.Sprintf("Error: Failed to deploy custom helm chart - %s", str)
	}
	return fmt.Sprintf("Profile %s deployed successfully", profileName)
}

func shellCommand(command, dir string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return string(out), err
}
