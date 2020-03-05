package gen

// Portal ...
type Portal struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

// Metadata ...
type Metadata struct {
	Name string `yaml:"metadata"`
}

// Spec ...
type Spec struct {
	Portal   string `yaml:"portal"`
	TargetIP string `yaml:"targetip"`
}
