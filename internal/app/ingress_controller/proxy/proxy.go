// This file sets up a reverse proxy for a kubernetes cluster
// automatically routing port 4000 to the port that the service
// on the kubernetes cluster is running on.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type Route struct {
	ServiceName string `json:"ServiceName"`
	ServicePort string `json:"ServicePort"`
	ServiceIP   string `json:"ServiceIP"`
}

type Rules struct {
	Protocol string `json:"Protocol"`
	Path     string `json:"Path"`
	Route    Route  `json:Route`
}

type Cluster struct {
	ClusterName string `json:"ClusterName"`
	ClusterIP   string `json:"ClusterIP"`
	ClusterPort string `json:"ClusterPort`
}

var rulesList = []Rules{}
var clusterList = []Cluster{}

// If anything is sent to the shutdown channel it will end the program.
var shutdownchan chan string = make(chan string)

func main() {
	fmt.Println("Software Defined Network Terminal")
	go GrabRules()                  // Constantly grab the servers
	go StartReverseProxy(PROXYPORT) // will set up the r-proxy when there is a server
	<-shutdownchan                  // wait here until there's a shutdown signal
	fmt.Println("Shuting Down...")
}

// GrabRules reads the file rules.json and puts the contents
// into the rulesList slice of rules struct
func GrabRules() {
	for {
		openFile, _ := ioutil.ReadFile("../rules.json")
		mu.Lock()
		err := json.Unmarshal(openFile, &rulesList)
		if err != nil {
			fmt.Println(err.Error())
		}
		mu.Unlock()
		time.Sleep(TIMETOSLEEP)
	}
}

// GrabClusters reads the file cluster.json and puts the onctents
// into the clusterList slice of clusters
func GrabClusters() {
	for {
		openFile, _ := ioutil.ReadFile("../clusters.json")
		mu.Lock()
		err := json.Unmarshal(openFile, &clusterList)
		if err != nil {
			fmt.Println(err.Error())
		}
		mu.Unlock()
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
		return
	}
	defer conn.Close()
	ConnSignal <- "New Connection \n"

	//Checking for server to handle the connecting client
	buf := make([]byte, 1024)
	conn.Read(buf)
	// parse for http or https and path
	request := string(buf)
	requestLine := strings.Split(request, "\n")[0]
	requestWords := strings.Split(requestLine, " ")
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
				conn.Write([]byte("Could not resolve: " + destination))
				fmt.Println("Could not resolve: " + destination)
				mu.Unlock()
				return
			}
			serverConn.Write(buf)
			break
		}
	}

	if serverConn == nil {
		for _, v := range clusterList {
			// send request to other clusters
			destination := v.ClusterIP + ":" + v.ClusterPort
			serverConn, err = net.Dial("tcp", destination)
			if err == nil {
				serverConn.Write(buf)
			}
		}
	}
	mu.Unlock()

	shutdownSession := make(chan string)
	if serverConn != nil {
		go SessionListener(serverConn, shutdownSession, conn)
		go SessionListener(conn, shutdownSession, serverConn)
		<-shutdownSession
	}

	conn.Write([]byte("404: Page not found on any cluster"))
}

//SessionListener listens for connections noise and sends it to the writer
func SessionListener(Conn net.Conn, shutdown chan string, Conn1 net.Conn) {
	var cnt = 0
	for {
		var temp []byte
		for {
			buf := make([]byte, 1024)
			Conn.SetReadDeadline(time.Now().Add(1 * time.Second))
			_, err := Conn.Read(buf)
			if err != nil {
				if len(temp) == 0 {
					cnt++
				} else {
					cnt = 0
				}
				break
			}
			temp = append(temp, buf...)
		}
		mu.Lock()
		Conn1.Write(temp)
		mu.Unlock()

		if cnt >= 50 {
			break
		}
	}
	shutdown <- "Ending"
}
