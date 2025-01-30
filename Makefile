# Define project-specific variables
PROJECT_NAME := mycontroller
IMAGE_NAME := mahdibouaziz/$(PROJECT_NAME)
VERSION := latest
GO_FILES := $(shell find . -name '*.go')

# Kubernetes config paths
CONFIG_DIR := config
CRD_DIR := $(CONFIG_DIR)/crd
RBAC_DIR := $(CONFIG_DIR)/rbac
MANAGER_DIR := $(CONFIG_DIR)/manager

# Tooling
CONTROLLER_GEN := controller-gen
KUBECTL := kubectl
DOCKER := docker
GO := go

# Default goal
.DEFAULT_GOAL := help

## 🛠️ Build the manager binary
build: $(GO_FILES)
	@echo "🔨 Building the manager binary..."
	$(GO) build -o bin/manager ./cmd/main.go

## 🧪 Run tests
test:
	@echo "🧪 Running tests..."
	$(GO) test ./... -cover

## 🎭 Generate CRDs, RBAC, and DeepCopy functions
generate:
	@echo "⚙️ Generating DeepCopy, CRDs, and RBAC files..."
	$(CONTROLLER_GEN) object paths=./api/... output:object:dir=./api/v1/
	$(CONTROLLER_GEN) crd paths=./api/... output:crd:dir=$(CRD_DIR)
	$(CONTROLLER_GEN) rbac:roleName=$(PROJECT_NAME) paths=./controllers/... output:rbac:dir=$(RBAC_DIR)

## 🐳 Build and push Docker image
docker-build:
	@echo "🐳 Building Docker image..."
	$(DOCKER) build -t $(IMAGE_NAME):$(VERSION) .

docker-push:
	@echo "🚀 Pushing Docker image..."
	$(DOCKER) push $(IMAGE_NAME):$(VERSION)

## 🚀 Deploy the CRD and controller
deploy-crd:
	@echo "📜 Applying CRDs..."
	$(KUBECTL) apply -f $(CRD_DIR)

deploy-rbac:
	@echo "🔒 Applying RBAC configurations..."
	$(KUBECTL) apply -f $(RBAC_DIR)

deploy-manager:
	@echo "🚀 Deploying the controller..."
	$(KUBECTL) apply -f $(MANAGER_DIR)

## 🚀 Deploy everything (CRD + RBAC + Manager)
deploy: deploy-crd deploy-rbac deploy-manager

## 🔄 Restart the controller
restart:
	@echo "♻️ Restarting the controller..."
	$(KUBECTL) rollout restart deployment/$(PROJECT_NAME)

## 📜 Check CRDs
get-crds:
	$(KUBECTL) get crds

## 📜 Check running Pods
get-pods:
	$(KUBECTL) get pods -l app=$(PROJECT_NAME)

## 🛠️ Display available Makefile commands
help:
	@echo "🛠️  Available Makefile Commands:"
	@echo "-------------------------------------------"
	@echo " build           - Build the manager binary"
	@echo " test            - Run tests"
	@echo " generate        - Generate DeepCopy, CRDs, and RBAC files"
	@echo " docker-build    - Build Docker image"
	@echo " docker-push     - Push Docker image"
	@echo " deploy-crd      - Deploy the CRDs"
	@echo " deploy-rbac     - Deploy RBAC permissions"
	@echo " deploy-manager  - Deploy the controller"
	@echo " deploy          - Deploy everything (CRD + RBAC + Controller)"
	@echo " restart         - Restart the controller"
	@echo " get-crds        - List installed CRDs"
	@echo " get-pods        - List running pods"