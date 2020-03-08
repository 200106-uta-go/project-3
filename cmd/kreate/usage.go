package main

const usage = `
mcproxy is a tool for creating and running reverse proxies.

Usage:

        mcproxy <command> [arguments]

The commands are:

		init		sets up necessary dependencies
		build		create a new mcProxy profile with a given name
		run 		start running mcProxy server with selected profile.
		profiles	list out reverse proxy profiles
		remove		remove selected profile from system
		mount		add a routing to an application server
		unmount		remove a routing to an application server
`
