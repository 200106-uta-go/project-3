package main

import "fmt"

func main() {
	// Set up error log to ouptut information
	log.SetFlags(log.Lshortfile)

	// Parse Command Line arguements and use them to find
	// appropriate sub-command for mcProxy to run.
	flag.Parse()
	switch flag.Arg(0) {
	case "build":
		proxy.CreateProfile(flag.Arg(1))
	case "mount":
		proxy.Mount(flag.Arg(1))
	case "run":
		proxy.Run(flag.Arg(1))
	case "unmount":
		proxy.UnMount(flag.Arg(1))
	case "profiles":
		proxy.ViewProfiles()
	case "remove":
		proxy.Remove(flag.Arg(1))
	case "help":
		fmt.Println(usage)
	default:
		fmt.Println("No valid sub command selected. Use \"mcproxy help\" for information on various options.")
	}
}
