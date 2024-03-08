
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLIST=$(GOCMD) list
BINARY_NAME=ksmapi
PKGS=$(shell ${GOLIST} ./... | grep -v /vendor)
COVERAGEDIR=./dist/coverage
TAG=$(shell git describe --tags --always --dirty)

# server
.PHONY: run
run:
	@echo "Make: Running app..."
	@$(GOCMD) run cmd/*.go --config configs/ksmapi.yaml ${ARGS}

.PHONY: test
test:
	@echo "Make: Running tests..."
	@$(GOTEST) ${OPTS} $(PKGS)

.PHONY: test-coverage
test-coverage:
	@echo "Running test with coverage..."
	@$(GOTEST) -json -covermode=atomic -coverpkg=./... $(PKGS)

.PHONY: build
build:
	@if [ ! -f 'dist/.ksmapi.yaml' ]; then cp configs/ksmapi.yaml dist/ksmapi.yaml; fi;
	CGO_ENABLED=0 $(GOBUILD) -ldflags="-extldflags=-static -X hexagonal.software/ksm-api/internal/version.Version=${TAG}" -o dist/$(BINARY_NAME) cmd/*.go

.PHONY: clean
clean:
	rm -rf dist/*

.PHONY: release-docker
release-docker:
	@if [ ! -f 'dist/ksmapi.yaml' ]; then cp configs/ksmapi.yaml dist/ksmapi.yaml; fi;
	DOCKER_BUILDKIT=1 docker build --target PROD -f scripts/docker/Dockerfile ${OPTS} --build-arg APPVERSION=${TAG} --tag ${REPO}:${TAG} .
	@echo ">>> Docker image built: ${REPO}:${TAG}"
