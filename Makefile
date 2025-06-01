# Vista Makefile
# Provides commands for building, testing, and running the Vista API server

# Variables
BINARY_NAME=vista
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")
VERSION=0.1.0

# Default port for the server
PORT=8080

.PHONY: all build clean run test test-verbose test-api fmt lint help

all: fmt build test

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) ./cmd/vista

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@go clean

# Run the server
run:
	@echo "Running server on port $(PORT)..."
	@go run ./cmd/vista/main.go -port $(PORT)

# Run tests
test:
	@echo "Running tests..."
	@go test ./... -count=1

# Run tests with verbose output
test-verbose:
	@echo "Running tests with verbose output..."
	@go test ./... -v -count=1

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test ./... -cover -count=1 -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

# Format code
fmt:
	@echo "Formatting code..."
	@gofmt -w $(GO_FILES)

# Lint code
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with:"; \
		echo "  brew install golangci-lint"; \
	fi

# Test API endpoints
test-api:
	@echo "Testing API endpoints..."
	@./scripts/test-api.sh

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Show help
help:
	@echo "Vista Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build         Build the binary"
	@echo "  make clean         Clean build artifacts"
	@echo "  make run           Run the server (default port: $(PORT))"
	@echo "  make run PORT=9000 Run the server on a custom port"
	@echo "  make test          Run all tests"
	@echo "  make test-verbose  Run tests with verbose output"
	@echo "  make test-coverage Run tests with coverage report"
	@echo "  make test-api      Test API endpoints (server must be running)"
	@echo "  make fmt           Format code"
	@echo "  make lint          Lint code"
	@echo "  make deps          Install dependencies"
	@echo "  make all           Format, build, and test"
	@echo "  make help          Show this help message"
