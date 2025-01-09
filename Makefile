.PHONY: all build test clean lint install fmt

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOFMT=$(GOCMD) fmt
BINARY_NAME=gmail-brita

# Build flags
LDFLAGS=-ldflags "-s -w"

all: test build

fmt:
	$(GOFMT) ./...

build: fmt
	$(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/main.go

test: fmt
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f *.xml

lint: fmt
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

install: build
	$(GOINSTALL) ./...

# Development helpers
run-example: fmt
	go run ./examples/main.go

# Create example filter XML
example-filter: fmt
	@mkdir -p examples/output
	go run ./cmd/main.go -config examples/filters.yaml -out examples/output/filters.xml
