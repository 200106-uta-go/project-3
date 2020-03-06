// This file sets up a reverse proxy for a kubernetes cluster
// automatically routing port 4000 to the port that the service
// on the kubernetes cluster is running on.

package main

import (
	"fmt"
	"github/200106-uta-go/project-3/internal/app/ingress_controller/scanner"
	"net"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex // mutex lock

// TIMETOSLEEP is how long the program waits in between checking for services
const TIMETOSLEEP = 10 * time.Second

// PROXYPORT is a string of the port number where the reverse proxy can be accessed
const PROXYPORT = "80"
const SECONDPORT = "8080"

var rulesList = []scanner.Rules{}
var clusterList = []scanner.AltCluster{}

// If anything is sent to the shutdown channel it will end the program.
var shutdownchan chan string = make(chan string)

func main() {
	fmt.Println("Software Defined Network Terminal")
	go GrabRules()                   // Constantly grab the servers
	go StartReverseProxy(PROXYPORT)  // will set up the r-proxy when there is a server
	go StartReverseProxy(SECONDPORT) // will set up the r-proxy when there is a server
	<-shutdownchan                   // wait here until there's a shutdown signal
	fmt.Println("Shuting Down...")
}

// GrabRules reads the file rules.json and puts the contents
// into the rulesList slice of rules struct
func GrabRules() {
	for {

		temp1, temp2 := scanner.Scan()

		mu.Lock()
		rulesList, clusterList = temp1, temp2
		mu.Unlock()

		fmt.Println(rulesList)
		fmt.Println(clusterList)

		time.Sleep(TIMETOSLEEP)
	}
}

//StartReverseProxy begins the hosting process for the
//client to server application
func StartReverseProxy(port string) {
	fmt.Println("Launching Software Defined Network...")

	// Listen on the PROXYPORT
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// If there is some error shut down the proxy
		fmt.Println(err)
		shutdownchan <- "Could Not Listen on Port"
		return
	}

	fmt.Println("Online - Now Listening On Port: " + port)

	// Create a channel for the connection signal that allows us to
	// wait for a new connection continuously
	ConnSignal := make(chan string)

	for {
		// For every new connection make a session then listen for a new connection
		go Session(ln, ConnSignal, port)
		fmt.Println(<-ConnSignal)
	}

}

//Session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func Session(ln net.Listener, ConnSignal chan string, port string) {
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err.Error())
		ConnSignal <- "Bad Connection \n"
		return
	}
	defer conn.Close()
	ConnSignal <- "New Connection \n"

	var IsFirst = true
	if port != PROXYPORT {
		IsFirst = false
	}

	//Checking for server to handle the connecting client
	buf := make([]byte, 1024)
	conn.Read(buf)
	// parse for http or https and path
	request := string(buf)
	requestLine := strings.Split(request, "\n")[0]
	requestWords := strings.Split(requestLine, " ")
	if len(requestWords) <= 1 {
		conn.Write(buf)
		return
	}
	path := requestWords[1]
	pathKeys := strings.Split(path, "?")
	path = pathKeys[0]
	pathFragments := strings.Split(path, "#")
	path = pathFragments[0]
	protocol := strings.ToLower(strings.Split(requestWords[2], "/")[0])
	fmt.Println("path = ", path)
	fmt.Println("protocol = ", protocol)

	var serverConn net.Conn = nil

	mu.Lock()
	for _, v := range rulesList {
		if v.Path == path && strings.ToLower(v.Protocol) == protocol {
			// route to location specified by the rule
			route := v.Route
			destination := route.ServiceIP + ":" + route.ServicePort
			fmt.Println("Going to Send Conn to: " + destination)
			serverConn, err = net.Dial("tcp", destination)
			if err != nil {
				mu.Unlock()
				fmt.Println("Could not resolve: " + destination + " server could not be dialed: " + err.Error())
				serverConn = nil
				break
			}
			defer serverConn.Close()
			serverConn.Write(buf)
			break
		}
	}
	shutdownSession := make(chan string)

	if serverConn == nil && IsFirst {
		for _, v := range clusterList {
			// send request to other clusters
			// destination := v.ClusterIP + ":" + v.ClusterPort
			destination := v.ClusterIP
			serverConn, err = net.Dial("tcp", destination)
			if err == nil {
				serverConn.Write(buf)
				buf2 := make([]byte, 1024)
				serverConn.SetReadDeadline(time.Now().Add(5 * time.Second))
				_, err := serverConn.Read(buf2)
				if err != nil {
					fmt.Println("404 No Response from portal server, could dial but not read")
				} else if !strings.Contains(string(buf2), "404") {
					var temp = []byte{}

					for _, b := range buf2 {
						if b != byte('\u0000') {
							temp = append(temp, b)
						}
					}

					conn.Write(temp)

					go SessionListener(serverConn, shutdownSession, conn)
					go SessionListener(conn, shutdownSession, serverConn)

					mu.Unlock()
					fmt.Println(<-shutdownSession)
					return
				}
				serverConn = nil
			}
		}
	}
	mu.Unlock()

	if serverConn != nil {
		go SessionListener(serverConn, shutdownSession, conn)
		go SessionListener(conn, shutdownSession, serverConn)
		fmt.Println(<-shutdownSession)
	} else {
		conn.Write([]byte("404: Page not found on any cluster"))
	}
}

//SessionListener listens for connections noise and sends it to the writer
func SessionListener(Conn net.Conn, shutdown chan string, Conn1 net.Conn) {
	for {

		buf := make([]byte, 1024)
		Conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		_, err := Conn.Read(buf)
		if err != nil {
			return
		}

		var temp = []byte{}

		for _, b := range buf {
			if b != byte('\u0000') {
				temp = append(temp, b)
			}
		}

		Conn1.Write(temp)
	}
}
