package ingressutil

//Ingress holds data for a Kubernets ingress deployment
type Ingress struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		TLS []struct {
			Hosts      []string `yaml:"hosts"`
			SecretName string   `yaml:"secretName"`
		} `yaml:"tls"`
		Rules []Rule `yaml:"rules"`
	} `yaml:"spec"`
}

//HTTP ...
type HTTP struct {
	Paths []Paths `yaml:"paths"`
}

//Rule ...
type Rule struct {
	Host string `yaml:"host"`
	HTTP HTTP   `yaml:"http"`
}

//Backend ...
type Backend struct {
	ServiceName string `yaml:"serviceName"`
	ServicePort int    `yaml:"servicePort"`
}

//Paths ...
type Paths struct {
	Path    string  `yaml:"path"`
	Backend Backend `yaml:"backend"`
}
