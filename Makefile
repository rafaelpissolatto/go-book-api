
# Makefile for building and pushing docker images

# Image URL to use all building/pushing image targets
IMAGE_NAME ?= rafaelpissolatto/api-devbook
IMAGE_VERSION ?= v0.0.1

# Build the docker image
build:
	echo "Building image $(IMAGE_NAME):$(IMAGE_VERSION)"
	docker build -t $(IMAGE_NAME) .
	docker tag $(IMAGE_NAME) $(IMAGE_NAME):$(IMAGE_VERSION)
	docker tag $(IMAGE_NAME) $(IMAGE_NAME):latest


# Push the docker image
push:
	echo "Pushing image $(IMAGE_NAME):$(IMAGE_VERSION)"
	docker push $(IMAGE_NAME):$(IMAGE_VERSION)
	docker push $(IMAGE_NAME):latest


# Helm upgrade the chart
upgrade:
	echo "Upgrading chart $(IMAGE_NAME):$(IMAGE_VERSION)"
	helm upgrade --install api-devbook deployments/helm/api-devbook


# Helm delete the chart
delete:
	echo "Deleting chart $(IMAGE_NAME):$(IMAGE_VERSION)"
	helm delete api-devbook


# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Run tests
test: fmt vet
	go test ./... -coverprofile cover.out

# kube port-forward
port-forward:
	kubectl port-forward svc/api-devbook 5000

## --------------------------------------
## Help
## --------------------------------------

.PHONY: help
help: Makefile
	@sed -n 's/^##//p' $<

# Path: Makefile