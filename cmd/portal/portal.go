package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"strings"
)

var apiPort = "12345"
var serverPort = "8081"

func main() {
	http.HandleFunc("/", router)

	//start up a kubectl proxy to locate api
	go startAPIProxy(apiPort)

	//convert to ListenAndServeTLS later
	fmt.Printf("Server is listening on localhost:%s \n", serverPort)
	log.Fatalln(http.ListenAndServe(":"+serverPort, nil))
}

func startAPIProxy(port string) {
	//build the command
	cmd := exec.Command("kubectl", "proxy", "--port="+port)

	//set stderr to output to terminal
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	cmd.Run()
	if stderr.String() != "" {
		log.Println(stderr.String())
	}
}

//router gets all portal addresses and sends the request to each foriegn cluster
func router(w http.ResponseWriter, r *http.Request) {

	portals := getportals("http://localhost:" + apiPort)

	//forward the request through each portal -- this needs to be re-written to check for successful reponse
	//currently if more than one portal get forwarded, the request gets more than one reponses
	for _, portal := range portals.Items {
		forwardRequest(w, r, portal.Spec.Targetip)
	}
}

//getPortals gets data for all portals from api server at the target address
func getportals(targetAddr string) Portals {
	//create address to get portals from kubernetes api
	// resp, err := http.Get("http://" + targetAddr + "/apis/revature.com/v1/namespaces/default/portals/")
	resp, err := http.Get(targetAddr + "/apis/revature.com/v1/namespaces/default/portals/")
	if err != nil {
		log.Fatalln(err)
	}

	//parse request body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//put body into portals struct and return it
	portals := Portals{}
	err = json.Unmarshal(body, &portals)
	if err != nil {
		log.Fatalln(err)
	}
	return portals
}

//forwardRequest
func forwardRequest(w http.ResponseWriter, r *http.Request, targetAddr string) {
	//add http protocol prefix if not already present
	if !strings.HasPrefix(targetAddr, "http://") && !strings.HasPrefix(targetAddr, "https://") {
		targetAddr = "http://" + targetAddr
	}

	//parse the string url into a url struct for reverse proxy
	url, err := url.Parse(targetAddr)
	if err != nil {
		log.Fatalln(err)
	}

	//setup reverse proxy to forward request to new host
	fmt.Println("Forwarding to " + url.String())
	forward := httputil.NewSingleHostReverseProxy(url)

	//update the headers to allow for SSL redirection
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	forward.ServeHTTP(w, r)
}
