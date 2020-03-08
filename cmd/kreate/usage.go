package main

const usage = `
kreate is a tool for setting up a kubernetes cluster with the ability to contact other
Kubernetes clusters in the event that a service can not be found. The tool will also set
up additional resources, such as Istio, Gafana, & Jaeger.

Usage:

        kreate <sub-command> [PROFILE_NAME]

The commands are:

		init		sets up necessary dependencies
		profile		create a new profile which the user can set key value pairs in "etc/kreate"
		edit		edit a pre-existing profile
		chart		will use a profile to produce a preconfigured chart that the user may use with his verstion of helm
		remove		remove selected profile from system
		ls  		will produce a list of available profiles
		run 		will apply the user defined profile to the Kubernetes cluster
		help		display this help text
`
