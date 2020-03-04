package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/200106-uta-go/project-3/pkg/kreate"
)

func main() {
	// Set up error log to ouptut information
	log.SetFlags(log.Lshortfile)

	// Must run Init() to set up environmental variables that contain path to our main data folder
	// that contains the helm chart directories. The environmenta
	kreate.Init()

	// Parse Command Line arguements and use them to find
	// appropriate sub-command for kreate to run.
	flag.Parse()

	switch flag.Arg(0) {
	case "init":
		// Sets up directory paths and environment variables
		kreate.Initalization()
	case "profile":
		// Creates a new profile .yaml
		kreate.CreateProfile(flag.Arg(1))
	case "edit":
		// Edits an existing profile .yaml
		kreate.EditProfile(flag.Arg(1))
	case "remove":
		// Removes an existing profile .yaml
		kreate.RemoveProfile(flag.Arg(1))
	case "chart":
		// Builds a helm chart using the specified profile
		kreate.CreateChart(flag.Arg(1))
	case "run":
		// Installs the istio environment (if not already installed) and Installs/Upgrades a helm chart using the specified profile
		kreate.RunProfile(flag.Arg(1))
	case "help":
		// describes CLI commands to user
		fmt.Println(usage)
		fallthrough
	default:
		fmt.Println("No valid sub command selected. Use \"kreate help\" for information on various options.")
	}
}
