package yaml

//Deployment ...
type Deployment struct {
	APIVersion string             `yaml:"apiVersion"`
	Kind       string             `yaml:"kind"`
	Metadata   ObjectMetaType     `yaml:"metadata"`
	Spec       DeploymentSpecType `yaml:"spec"`
}

//DeploymentSpecType ...
type DeploymentSpecType struct {
	MinReadySeconds         int  `yaml:"minReadySeconds"`
	Paused                  bool `yaml:"paused"`
	ProgressDeadlineSeconds int  `yaml:"progressDeadlineSeconds"`
	Replicas                int  `yaml:"replicas"`
	RevisionHistoryLimit    int  `yaml:"revisionHistoryLimit"`
	Selector                LabelSelectorType
	Strategy                DeploymentStrategyType
	Template                PodTemplateSpecType
}

//DeploymentStrategyType ...
type DeploymentStrategyType struct {
	RollingUpdate RollingUpdateDeploymentType `yaml:"rollingUpdate"`
	Type          string                      `yaml:"type"`
}

//RollingUpdateDeploymentType ...
type RollingUpdateDeploymentType struct {
	maxSurge       int `yaml:"maxSurge"`
	maxUnavailable int `yaml:"maxUnavailable"`
}
