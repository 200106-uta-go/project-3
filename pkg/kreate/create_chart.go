package kreate

import "os/exec"

// import the helm code here to implement the function

// CreateChart will take in a path to a set up yaml, and then create the output chart in the user's current directory.
func CreateChart(name string) error {
	// Use Helm code here to create charts for MVP

	exec.Command("helm", "create", name)

	return nil
}
