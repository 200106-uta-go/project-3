package main

//Portals holds all data returned from the Kubernetes API when getting all portals in a cluster
type Portals struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			CreationTimestamp string `json:"creationTimestamp"`
			Generation        int    `json:"generation"`
			Name              string `json:"name"`
			Namespace         string `json:"namespace"`
			ResourceVersion   string `json:"resourceVersion"`
			SelfLink          string `json:"selfLink"`
			UID               string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			Cluster  string `json:"cluster"`
			Targetip string `json:"targetip"`
		} `json:"spec"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		Continue        string `json:"continue"`
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
	} `json:"metadata"`
}
