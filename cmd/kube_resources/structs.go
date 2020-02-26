package main

//Pipes holds all data returned from the Kubernetes API when getting all pipes in a cluster
type Pipes struct {
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

//Services holds all data returned from the Kubernetes API when getting all services in a cluster
type Services struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		Metadata struct {
			CreationTimestamp string `json:"creationTimestamp"`
			Labels            struct {
				Component string `json:"component"`
				Provider  string `json:"provider"`
			} `json:"labels"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			ResourceVersion string `json:"resourceVersion"`
			SelfLink        string `json:"selfLink"`
			UID             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			ClusterIP string `json:"clusterIP"`
			Ports     []struct {
				Name       string `json:"name"`
				Port       int    `json:"port"`
				Protocol   string `json:"protocol"`
				TargetPort int    `json:"targetPort"`
			} `json:"ports"`
			SessionAffinity string `json:"sessionAffinity"`
			Type            string `json:"type"`
		} `json:"spec"`
		Status struct {
			LoadBalancer struct{} `json:"loadBalancer"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
	} `json:"metadata"`
}
