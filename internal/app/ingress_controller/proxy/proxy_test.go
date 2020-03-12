package main

import (
	"net"
	"testing"
	"time"
)

//TestSession ensures that we can establish a connection on port 4000
//and dial to this connection without an error
func TestSession(t *testing.T) {
	myPort := "4000"
	go StartReverseProxy(myPort)
	time.Sleep(3 * time.Second)

	_, err := net.Dial("tcp", ":"+myPort)

	if err != nil {
		t.Error(err)
	}
}
