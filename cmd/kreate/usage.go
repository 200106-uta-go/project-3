package main

const usage = `
mcproxy is a tool for creating and running reverse proxies.

Usage:

        mcproxy <command> [arguments]

The commands are:

		init		sets up necessary dependencies
		build		create a new mcProxy profile with a given name
		run 		start running mcProxy server with selected profile
		profile		create a new profile under "/etc/kreate/" directory
		remove		remove selected profile from system
		mount		add a routing to an application server
		unmount		remove a routing to an application server
`

// profiles	list out reverse proxy profiles - removed because is not a used command
