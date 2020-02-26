package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/", router)
	http.HandleFunc("/pipes", pipeHandler)
	http.HandleFunc("/services", serviceHandler)

	//convert to ListenAndServeTLS later
	fmt.Println("Server is listening on localhost:8081")
	log.Fatalln(http.ListenAndServe(":8081", nil))
}

func router(w http.ResponseWriter, r *http.Request) {
	//check what the intended svc/pod is
	targetSvc := r.Header.Get("targetSvc")
	if targetSvc == "" {
		log.Println("targetSvc header not present in request")
	}

	pipes := getPipes("http://127.0.0.1:8080")

	//find the address of the cluster containing the target service
	targetAddr := findService(targetSvc, pipes)

	//forward data to the next destination
	forwardRequest(w, r, targetAddr)
}

//handler to return all pipes for the local cluster
func pipeHandler(w http.ResponseWriter, r *http.Request) {
	pipes := getPipes("http://127.0.0.1:8080")

	bytes, err := json.Marshal(pipes)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(bytes)
}

//handler to return all services for the local cluster
func serviceHandler(w http.ResponseWriter, r *http.Request) {
	services := getServices("http://127.0.0.1:8080")

	bytes, err := json.Marshal(services)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(bytes)
}

//get data for all pipes from api server
func getPipes(targetAddr string) Pipes {
	//create address to get pipes from kubernetes api
	resp, err := http.Get(targetAddr + "/apis/revature.com/v1/namespaces/default/pipes/")
	if err != nil {
		log.Fatalln(err)
	}

	//parse request body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//put body into pipes struct and return it
	pipes := Pipes{}
	err = json.Unmarshal(body, &pipes)
	if err != nil {
		log.Fatalln(err)
	}
	return pipes
}

func getServices(targetAddr string) Services {
	//create address to get pipes from kubernetes api
	resp, err := http.Get(targetAddr + "/api/v1/namespaces/default/services")
	if err != nil {
		log.Fatalln(err)
	}

	//parse request body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//put body into pipes struct and return it
	services := Services{}
	err = json.Unmarshal(body, &services)
	if err != nil {
		log.Fatalln(err)
	}
	return services
}

func findService(target string, pipes Pipes) string {
	serviceAddr := ""

	//get all services in each cluster
	for _, pipe := range pipes.Items {
		services := getServices(pipe.Spec.Targetip)

		//for each service in a cluster, check if it matches the target
		for _, svc := range services.Items {
			if svc.Metadata.Name == target {
				serviceAddr = svc.Spec.ClusterIP
			}
		}
	}

	//search the struct for all matches for svc
	return serviceAddr
}

func forwardRequest(w http.ResponseWriter, r *http.Request, targetAddr string) {
	//parse the string url into a url struct for reverse proxy
	url, err := url.Parse(targetAddr)
	if err != nil {
		log.Fatalln(err)
	}

	//setup reverse proxy to forward request to new host
	forward := httputil.NewSingleHostReverseProxy(url)

	//update the headers to allow for SSL redirection
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	forward.ServeHTTP(w, r)
}
