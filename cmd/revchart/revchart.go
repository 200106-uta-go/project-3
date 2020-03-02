package main

import (
	"flag"
	"fmt"

	ch "github.com/200106-uta-go/project-3/pkg/chart"
	gen "github.com/200106-uta-go/project-3/pkg/create"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "create":
		c := ch.Metadata{
			Name: flag.Arg(1),
		}
		gen.Create(&c, "./")
	default:
		fmt.Println("No valid sub command selected. Use \"revchart help\" for information on various options.")
	}
}
