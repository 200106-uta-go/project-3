package yamlHandler

import "time"

type Metadata struct {
	Annotation                 string    `yaml:"annotation"`
	ClusterName                string    `yaml:"clusterName"`
	CreationTimestamp          time.Time `yaml:"creationTimestamp"`
	DeletionGracePeriodSeconds int       `yaml:"deletionGracePeriodSeconds"`
	DeletionTimestamp          time.Time `yaml:"deletionTimestamp"`
	Finalizers                 []string  `yaml:"finalizers"`
	GenerateName               string    `yaml:"generateName"`
	Generation                 int       `yaml:"generation"`
	Initializers                         // Need initializer type to be defined
	labels                               // string??
	ManagedFields                        // ManagedFieldEntry array
	Name                       string    `yaml:"name"`
	Namespace                  string    `yaml:"namespace`
	OwnerReferences                      // OwnerReferences type to be defined
	ResourceVersion            string    `yaml:"resourceVersion"`
	SelfLink                   string    `yaml:"selfLink"`
	Uid                        string    `yaml:"uid"`
}
