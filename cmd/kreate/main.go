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

	// Parse Command Line arguements and use them to find
	// appropriate sub-command for mcProxy to run.
	flag.Parse()
	switch flag.Arg(0) {
	case "chart":
		kreate.CreateChart(flag.Arg(1))

	//SCAFFOLDING FOR AFTER MVP

	// case "mount":
	// 	kreate.Mount(flag.Arg(1))
	// case "run":
	// 	kreate.Run(flag.Arg(1))
	// case "unmount":
	// 	kreate.UnMount(flag.Arg(1))
	// case "profiles":
	// 	kreate.ViewProfiles()
	// case "remove":
	// 	kreate.Remove(flag.Arg(1))

	case "help":
		fmt.Println(usage)
	default:
		fmt.Println("No valid sub command selected. Use \"mckreate.help\" for information on various options.")
	}
}
