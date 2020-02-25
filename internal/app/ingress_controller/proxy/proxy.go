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

var mu sync.Mutex

const TIMETOSLEEP = 10 * time.Second

//Power is a control bool to be accessed to shut down the
//clientserver
var backendServers map[string]string = make(map[string]string)
var shutdownchan chan string = make(chan string)

func main() {
	fmt.Println("Software Defined Network Terminal")
	go GrabServers()
	go StartReverseProxy("4000")
	<-shutdownchan
	fmt.Println("Shuting Down...")
}

//StartReverseProxy begins the hosting process for the
//client to server application
func StartReverseProxy(port string) {
	fmt.Println("Launching Software Defined Network...")

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		shutdownchan <- "Could Not Listen on Port"
		return
	}

	fmt.Println("Online - Now Listening On Port: " + port)

	ConnSignal := make(chan string)

	for {

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
					return
				} else {
					serverConn.Write(buf)
					break
				}
			}
		}
		if serverConn != nil {
			break
		}

	}
	mu.Unlock()

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
