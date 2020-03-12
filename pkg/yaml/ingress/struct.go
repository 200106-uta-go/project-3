package ingress

//Ingress holds data used to populate a kubernetes ingress yaml file
type Ingress struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

//Metadata ...
type Metadata struct {
	Name        string   `yaml:"name"`
	Labels      struct{} `yaml:",omitempty"`
	Annotations struct{} `yaml:",omitempty"`
}

//Spec ...
type Spec struct {
	TLS     []TLS   `yaml:",omitempty"`
	Rules   []Rules `yaml:",omitempty"`
	Backend Backend `yaml:",omitempty"`
}

//TLS ...
type TLS struct {
	Hosts      []string `yaml:"hosts"`
	SecretName string   `yaml:"secretName"`
}

//Rules ...
type Rules struct {
	Host string `yaml:"host"`
	HTTP HTTP   `yaml:"http"`
}

//HTTP ...
type HTTP struct {
	Paths []Paths `yaml:"paths"`
}

//Paths ...
type Paths struct {
	Path    string  `yaml:"path"`
	Backend Backend `yaml:"backend"`
}

//Backend ...
type Backend struct {
	ServiceName string `yaml:"serviceName"`
	ServicePort int    `yaml:"servicePort"`
}
