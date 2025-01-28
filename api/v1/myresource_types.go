// +groupName=sample.example.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MyResourceSpec defines the desired state of MyResource
type MyResourceSpec struct {
	ReplicaCount int    `json:"replicaCount"` // Number of replicas
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
