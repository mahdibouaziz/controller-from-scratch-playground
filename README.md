# controller-from-scratch-playground

# Objective

We aim to:

1. Define a Custom Resource Definition (**CRD**) manually.
2. Write a **controller** to **reconcile** the state of this custom resource.
3. Use tools like `controller-gen` and `controller-runtime` without relying on Kubebuilder.
4. Prepare everything (manifests, Makefile, and code) manually to understand the entire process in-depth.

# Steps

## Step 1: Setting Up the Project

Start by creating the directory structure for the project.

```bash
mkdir controller-from-scratch-playground
cd controller-from-scratch-playground
go mod init github.com/mahdibouaziz/controller-from-scratch-playground
```

This initializes a Go module. The `go.mod` file will track dependencies for the project.

## Step 2: Install Required Tools

Install the tools and libraries you'll need:

1. `controller-gen`: For generating CRDs, RBAC rules, and webhook configurations.
2. `controller-runtime`: A library to simplify writing controllers.
3. `client-go`: Kubernetes client library for interacting with the Kubernetes API.

```bash
# Install controller-gen
go install sigs.k8s.io/controller-tools/cmd/controller-gen@latest
cp ~/go/bin/controller-gen /usr/local/bin

# Add necessary libraries to your project
go get sigs.k8s.io/controller-runtime@latest
go get k8s.io/client-go@latest
```

## Step 3: Define the CRD

The Custom Resource Definition (CRD) specifies the schema for your custom Kubernetes resource.

1. Create the Go API definition: Create a file named `api/v1/myresource_types.go`
    - `+groupName` marker defines the API group for your CRD. Place it at the top of your myresource_types.go file: `// +groupName=sample.example.com`
    - `+kubebuilder:object:root=true`: Marks the type as the root object for the CRD. TODO - explain more this
    - `+kubebuilder:subresource:status`: Adds a status subresource for the CRD.

First create the required directories

```bash
mkdir api # This will contains the definition of your crds (golang types). you should respect the version v1alpha1, v1alpha2, v1beta1, v1beta2, v1.
mkdir api/v1 # for this example we are going directly to v1
```

```golang
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
```

2. Generate the CRD YAML: Use controller-gen to generate the CRD YAML file.
```bash
mkdir config # This will contains K8s configuration
mkdir config/crd # This will contains the generated crds by controller-gen

controller-gen crd paths=./api/... output:crd:dir=./config/crd
```
