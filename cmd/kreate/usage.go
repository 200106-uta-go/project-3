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
			-name	Sets the name of profile
			-clustername	Sets the clustername of the profile
			-clusterip	Sets the clusterip of the profile
			-clusterport	Append a clusterport to the profile
			
			-NameOfApp	Specifies the name of the app which will be modified by the app-related input flags
			-imageurl	An App-related flag. Sets the imageurl of the App specified by the NameOfApp flag
			-servicename	An App-related flag. Sets the servicename of the App specified by the NameOfApp flag
			-serviceport	An App-related flag. Sets the serviceport of the App specified by the NameOfApp flag
			-port	An App-related flag. Appends a port to the App specified by the NameOfApp flag
			-endpoint	An App-related flag. Appends an endpoint to the App specified by the NameOfApp flag	
		chart		will use a profile to produce a preconfigured chart that the user may use with his verstion of helm
		remove		remove selected profile from system
			-a, --all	removes all profiles from system
		ls  		will produce a list of available profiles
		run 		will apply the user defined profile to the Kubernetes cluster
		help		display this help text
`
