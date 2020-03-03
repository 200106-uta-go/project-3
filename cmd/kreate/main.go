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

	case "chart":
		kreate.CreateChart(flag.Arg(1))
	case "edit":
		kreate.EditValues(flag.Arg(1))
	case "help":
		fmt.Println(usage)
	default:
		fmt.Println("No valid sub command selected. Use \"kreate help\" for information on various options.")
	}
}
