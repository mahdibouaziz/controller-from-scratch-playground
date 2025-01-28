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

First create the required directories

```bash
mkdir api # This will contains the definition of your crds (respect the version v1alpha1, v1alpha2, v1beta1, v1beta2, v1)
mkdir api/v1 # for this example we are going directly to v1

mkdir config # This will contains K8s configuration
mkdir config/crd # This will contains the generated crds by controller-gen
```

1. Create the Go API definition: Create a file named `api/v1/myresource_types.go`
    - `+kubebuilder:object:root=true`: Marks the type as the root object for the CRD. TODO - explain more this
    - `+kubebuilder:subresource:status`: Adds a status subresource for the CRD.
2. Generate the CRD YAML: Use controller-gen to generate the CRD YAML file.
```bash
controller-gen crd paths=./api/... output:crd:dir=./config/crd
```
