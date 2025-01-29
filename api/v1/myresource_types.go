// +kubebuilder:object:generate=true
// +groupName=sample.example.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Define API Group and Version
var (
	GroupVersion = schema.GroupVersion{Group: "sample.example.com", Version: "v1"}

	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// MyResourceSpec defines the desired state of MyResource
type MyResourceSpec struct {
	ReplicaCount int32  `json:"replicaCount"` // Number of replicas
	Message      string `json:"message"`      // Message to display
}

// MyResourceStatus defines the observed state of MyResource
type MyResourceStatus struct {
	AvailableReplicas int `json:"availableReplicas"` // Current replicas available
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MyResource represents the CRD
type MyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyResourceSpec   `json:"spec,omitempty"`
	Status MyResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MyResourceList is a list of MyResource
type MyResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyResource `json:"items"`
}

// Register MyResource and MyResourceList with the scheme
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&MyResource{},
		&MyResourceList{},
	)
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
