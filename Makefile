BINARY = blackbox-helloworld-responder
COMMIT := $(shell git rev-parse HEAD)
BRANCH := $(shell git symbolic-ref --short -q HEAD || echo HEAD)
DATE := $(shell date -u +%Y%m%d-%H:%M:%S)
VERSION_PKG = github.com/KohlsTechnology/blackbox-helloworld-responder/pkg/version
LDFLAGS := "-X ${VERSION_PKG}.Branch=${BRANCH} -X ${VERSION_PKG}.BuildDate=${DATE} \
	-X ${VERSION_PKG}.GitSHA1=${COMMIT}"
TAG?=""

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -rf $(BINARY) dist/

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(BINARY) -ldflags $(LDFLAGS)

.PHONY: image
image:
	docker build . -t quay.io/kohlstechnology/blackbox-helloworld-responder:latest

.PHONY: test
test: lint-all test-unit

.PHONY: test-unit
test-unit:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

# Make sure go.mod and go.sum are not modified
.PHONY: test-dirty
test-dirty: build
	git diff --exit-code

# Make sure goreleaser is working
.PHONY: test-release
test-release:
	BRANCH=$(BRANCH) COMMIT=$(COMMIT) DATE=$(DATE) VERSION_PKG=$(VERSION_PKG) goreleaser --snapshot --skip-publish --rm-dist

.PHONY: lint
lint:
	LINT_INPUT="$(shell go list ./...)"; golint -set_exit_status $$LINT_INPUT

.PHONY: golangci-lint
golangci-lint:
	golangci-lint run

.PHONY: lint-all
lint-all: lint golangci-lint

.PHONY: tag
tag:
	git tag -a $(TAG) -m "Release $(TAG)"
	git push origin $(TAG)

# Requires GITHUB_TOKEN environment variable to be set
.PHONY: release
release:
	BRANCH=$(BRANCH) COMMIT=$(COMMIT) DATE=$(DATE) VERSION_PKG=$(VERSION_PKG) goreleaser
