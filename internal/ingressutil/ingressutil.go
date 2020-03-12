package main

import (
	"os"

	"github.com/200106-uta-go/project-3/pkg/ingressutil"
)

func main() {
	host := "google.com"
	path := "/google"
	svcName := "zach-server"
	svcPort := 4000

	file, err := os.Open("./deployments/kubernetes/ingress.yaml")
	if err != nil {
		panic(err)
	}

	ingressutil.AddPath(file, host, path, svcName, svcPort)
}
