package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	go commandListener("8080")

	readCommands()
}

func readCommands() {
	for {
		mu.Lock()
		commandFile, _ := os.OpenFile("../commandlist", os.O_RDWR|os.O_CREATE, 7777)
		AllBytes, _ := ioutil.ReadAll(commandFile)
		commandFile.Truncate(0)
		mu.Unlock()

		var temp []byte = nil
		for _, v := range AllBytes {
			if v != 0 {
				temp = append(temp, v)
			}
		}
		AllBytes = temp

		Lines := strings.Split(string(AllBytes), "\n")

		for k, v := range Lines {
			command := strings.Split(string(v), " ")
			exec.Command(command[0], command[1:]...).Output()

			if len(command) >= 2 && command[1] == "expose" {
				output, _ := exec.Command("kubectl", "get", "svc").Output()

				t := strings.Split(string(output), "\n")
				t = t[1:]

				for _, v := range t {
					z := strings.Split(v, " ")

					var temp []string

					for k2, v2 := range z {
						z[k2] = strings.TrimSpace(v2)
						if z[k2] != "" {
							temp = append(temp, z[k2])
						}
					}

					z = temp

					if len(z) >= 2 && z[0] == command[3] {
						z = strings.Split(z[4], "/")
						l := strings.Split(z[0], ":")

						openFile, _ := ioutil.ReadFile("../serverlist.json")
						maps := make(map[string]string)
						json.Unmarshal(openFile, &maps)

						temp := Lines[k+1]

						maps[string(temp)] = l[1]

						bytestowrite, _ := json.MarshalIndent(maps, "", "	")

						ioutil.WriteFile("../serverlist.json", bytestowrite, 7777)

					}

				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func commandListener(port string) {
	//Pipe to control connection flow
	commandFile, _ := os.OpenFile("../commandlist", os.O_RDWR|os.O_CREATE, 7777)
	conPipe := make(chan string)

	ls, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Printf("Listen Error: %s", err)
	}

	defer ls.Close()

	//Cycle for loop as connections are made
	for {
		go commandConnection(ls, conPipe, commandFile)
		<-conPipe
	}
}

func commandConnection(ls net.Listener, conPipe chan string, commandFile *os.File) {
	con, err := ls.Accept()
	defer con.Close()
	conPipe <- "Connection made"
	if err != nil {
		fmt.Printf("Accept error: %s", err)
		return
	}

	buf := make([]byte, 1024)
	con.Read(buf)

	var temp []byte = nil
	for _, v := range buf {
		if v != 0 {
			temp = append(temp, v)
		}
	}
	buf = temp

	mu.Lock()
	commandFile.Write(temp)
	mu.Unlock()
}
