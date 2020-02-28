package yaml

import "time"

// ObjectMetaType is metadata that all persisted resources must have, which includes all objects users must create.
type ObjectMetaType struct {
	Annotations                map[string]string        `yaml:"annotations"`
	ClusterName                string                   `yaml:"clusterName"`
	CreationTimestamp          time.Time                `yaml:"creationTimestamp"`          // Read-only
	DeletionGracePeriodSeconds int                      `yaml:"deletionGracePeriodSeconds"` // Read-only
	DeletionTimestamp          time.Time                `yaml:"deletionTimestamp"`          // Read-only
	Finalizers                 []string                 `yaml:"finalizers"`
	GenerateName               string                   `yaml:"generateName"`
	Generation                 int                      `yaml:"generation"`
	Initializers               InitializersType         `yaml:"initializers"`
	Labels                     map[string]string        `yaml:"labels"`
	ManagedFields              []ManagedFieldsEntryType `yaml:"managedFields"`
	Name                       string                   `yaml:"name"`
	Namespace                  string                   `yaml:"namespace"`
	OwnerReferences            []OwnerReferenceType     `yaml:"ownerReferences"`
	ResourceVersion            string                   `yaml:"resourceVersion"`
	SelfLink                   string                   `yaml:"selfLink"`
	UID                        string                   `yaml:"uid"`
}

// InitializersType tracks the progress of initialization.
type InitializersType struct {
	Pending []InitializerType `yaml:"Pending"`
	Result  StatusType        `yaml:"result"`
}

// InitializerType is information about an initializer that has not yet completed.
type InitializerType struct {
	Name string `yaml:"name"`
}

// StatusType is a return value for calls that don't return other objects.
type StatusType struct {
	APIVersion string            `yaml:"apiVersion"`
	Code       int               `yaml:"code"`
	Details    StatusDetailsType `yaml:"details"`
	Kind       string            `yaml:"kind"`
	Message    string            `yaml:"message"`
	Metadata   ListMetaType      `yaml:"metadata"`
	Reason     string            `yaml:"reason"`
	Status     string            `yaml:"status"`
}

// StatusDetailsType is a set of additional properties that MAY be set by the server to provide additional information about a response.
// The Reason field of a Status object defines what attributes will be set.
// Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
type StatusDetailsType struct {
	Causes            []StatusCauseType `yaml:"causes"`
	Group             string            `yaml:"group"`
	Kind              string            `yaml:"kind"`
	Name              string            `yaml:"Name"`
	RetryAfterSeconds int               `yaml:"retryAfterSeconds"`
	UID               string            `yaml:"uid"`
}

// StatusCauseType provides more information about an api.Status failure, including cases when multiple errors are encountered.
type StatusCauseType struct {
	Field   string `yaml:"field"`
	Message string `yaml:"message"`
	Reason  string `yaml:"reason"`
}

// ListMetaType describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
type ListMetaType struct {
	Continue        string `yaml:"continue"`
	ResourceVersion string `yaml:"resourceVersion"`
	SelfLink        string `yaml:"selfLink"`
}

// ManagedFieldsEntryType is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
type ManagedFieldsEntryType struct {
	APIVersion string            `yaml:"apiVersion"`
	Fields     map[string]string `yaml:"field"`
	Manager    string            `yaml:"manager"`
	Operation  string            `yaml:"operation"`
	Time       time.Time         `yaml:"time"`
}

// OwnerReferenceType contains enough information to let you identify an owning object.
// An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.
type OwnerReferenceType struct {
	APIVersion         string `yaml:"apiVersion"`
	BlockOwnerDeletion bool   `yaml:"blockOwnerDeletion"`
	Controller         bool   `yaml:"controller"`
	Kind               string `yaml:"kind"`
	Name               string `yaml:"name"`
	UID                string `yaml:"uid"`
}

// LabelSelectorType is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed.
// An empty label selector matches all objects. A null label selector matches no objects.
type LabelSelectorType struct {
	MatchExpressions []LabelSelectorRequirementType `yaml:"LabelSelectorRequirement"`
	MatchLabels      map[string]string              `yaml:"matchLabels"`
}

// LabelSelectorRequirementType is a selector that contains values, a key, and an operator that relates the key and values.
type LabelSelectorRequirementType struct {
	Key      string `yaml:"key"`
	Operator string `yaml:"operator"`
	Values   string `yaml:"values"`
}
