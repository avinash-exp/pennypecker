BINARY=pennypecker
VERSION=0.1.0-alpha.1
MODULE=github.com/avinash-exp/pennypecker
GOLANGCI_LINT_VERSION=v2.7.2

.PHONY: build clean test fmt lint tidy init install-tools

default: build

install-tools:
	@echo "==> Installing development tools..."
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

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
	go test -v ./tests/...

fmt:
	@echo "==> Formatting code..."
	go fmt ./...

lint:
	@echo "==> Running linter..."
	golangci-lint run ./...

tidy:
	@echo "==> Tidying modules..."
	go mod tidy
