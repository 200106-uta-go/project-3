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
const PROXYPORT = "4000"

// The list of backend servers running on the kubernetes cluster pulled from JSON
var backendServers map[string]string = make(map[string]string)

// If anything is sent to the shutdown channel it will end the program.
var shutdownchan chan string = make(chan string)

func main() {
	fmt.Println("Software Defined Network Terminal")
	go GrabServers()                // Constantly grab the servers
	go StartReverseProxy(PROXYPORT) // will set up the r-proxy when there is a server
	<-shutdownchan                  // wait here until there's a shutdown signal
	fmt.Println("Shuting Down...")
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
		<-ConnSignal

	}

}

//Session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func Session(ln net.Listener, ConnSignal chan string, port string) {
	conn, _ := ln.Accept()
	defer conn.Close()
	ConnSignal <- "New Connection \n"

	//Checking for server to handle the connecting client
	buf := make([]byte, 1024)
	conn.Read(buf)
	var serverConn net.Conn = nil
	var err error
	for {
		mu.Lock()
		for k, v := range backendServers {
			if strings.Contains(string(buf), k) {
				serverConn, err = net.Dial("tcp", ":"+v)
				if err != nil {
					conn.Write([]byte("Could not resolve: " + ":" + v))
					fmt.Println("Could not resolve: " + ":" + v)
					mu.Unlock()
					return
				} else {
					serverConn.Write(buf)
					break
				}
			}
		}
		if serverConn != nil {
			mu.Unlock()
			break
		}

	}

	shutdownSession := make(chan string)
	go SessionListener(serverConn, shutdownSession, conn)
	go SessionListener(conn, shutdownSession, serverConn)
	<-shutdownSession
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

//GrabServers test
func GrabServers() {
	for {

		openFile, _ := ioutil.ReadFile("../serverlist.json")

		mu.Lock()
		_ = json.Unmarshal(openFile, &backendServers)
		mu.Unlock()

		time.Sleep(TIMETOSLEEP)
	}
}
