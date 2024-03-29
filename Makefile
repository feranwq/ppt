BINDIR := $(PWD)/bin
OUTPUT_DIR := $(PWD)/_output

GOOS ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
GOARCH ?= amd64

PATH := $(BINDIR):$(PATH)
SHELL := env PATH='$(PATH)' /bin/sh

all: build

# Run tests
test: fmt vet
	@# Disable --race until https://github.com/kubernetes-sigs/controller-runtime/issues/1171 is fixed.
	ginkgo --randomizeAllSpecs --randomizeSuites --failOnPending --flakeAttempts=2 \
			--cover --coverprofile cover.out --trace --progress  $(TEST_ARGS)\
			./pkg/... .

# Build ptt binary
build: fmt vet
	go build -o $(OUTPUT_DIR)/ppt .

# Cross compiler
build-all: fmt vet
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o $(OUTPUT_DIR)/ppt_amd64.exe .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o $(OUTPUT_DIR)/ppt_linux_amd64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o $(OUTPUT_DIR)/ppt_darwin_amd64 .

# Run go fmt against code
fmt:
	go fmt ./pkg/... .

# Run go vet against code
vet:
	go vet ./pkg/... .

lint:
	$(BINDIR)/golangci-lint run --timeout 2m0s ./pkg/... .

dependencies:
	test -d $(BINDIR) || mkdir $(BINDIR)
	GOBIN=$(BINDIR) go install github.com/onsi/ginkgo/ginkgo@v1.16.4
	
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.46.2