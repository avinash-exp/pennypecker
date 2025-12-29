BINARY=pennypecker
VERSION=0.1.0-alpha.1
MODULE=github.com/avinash-exp/pennypecker

.PHONY: build clean test fmt lint tidy init

default: build

init:
	@echo "==> Initializing Go module..."
	@if [ ! -f go.mod ]; then go mod init $(MODULE); fi
	go mod tidy

build:
	@echo "==> Building provider..."
	go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY)

clean:
	@echo "==> Cleaning..."
	rm -f $(BINARY)

test:
	@echo "==> Running tests..."
	go test -v ./...

fmt:
	@echo "==> Formatting code..."
	go fmt ./...

lint:
	@echo "==> Running linter..."
	golangci-lint run ./...

tidy:
	@echo "==> Tidying modules..."
	go mod tidy
