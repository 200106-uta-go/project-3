package kreate

// Profile defines the profile struct and converts between the struct format and the yaml format
type Profile struct {
	Name         string
	ClusterName  string
	ClusterIP    string
	ClusterPorts []string
	Apps         []App
}

// App ...
type App struct {
	Name      string
	ImageURL  string
	Ports     []string
	Endpoints []string
}
