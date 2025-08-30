# Proton Makefile

.PHONY: help build test clean install lint fmt vet check-deps docs

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
build: ## Build the proton binary
	@echo "Building proton..."
	go build -o bin/proton ./cmd/proton

build-all: ## Build for all platforms
	@echo "Building for all platforms..."
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/proton-linux-amd64 ./cmd/proton
	GOOS=darwin GOARCH=amd64 go build -o dist/proton-darwin-amd64 ./cmd/proton
	GOOS=darwin GOARCH=arm64 go build -o dist/proton-darwin-arm64 ./cmd/proton
	GOOS=windows GOARCH=amd64 go build -o dist/proton-windows-amd64.exe ./cmd/proton

install: ## Install proton to GOPATH/bin
	@echo "Installing proton..."
	go install ./cmd/proton

# Development targets
test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -cover ./...
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

# Quality checks
check: fmt vet lint test ## Run all checks

check-deps: ## Check for outdated dependencies
	@echo "Checking dependencies..."
	go list -u -m all

# Utility targets
clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/ dist/ coverage.out coverage.html

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download

tidy: ## Tidy go modules
	@echo "Tidying go modules..."
	go mod tidy

# Documentation targets
docs: ## Generate documentation for this project using proton
	@echo "Generating documentation..."
	@if [ ! -f bin/proton ]; then make build; fi
	./bin/proton generate

docs-serve: ## Serve documentation locally (requires gitbook-cli)
	@echo "Serving documentation..."
	@cd docs && gitbook serve

# Development helpers
dev-setup: ## Set up development environment
	@echo "Setting up development environment..."
	go mod download
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "Installing golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2; \
	fi

example: ## Run proton on example project
	@echo "Running proton on example..."
	@mkdir -p example-output
	./bin/proton generate --output example-output

# Release targets
tag: ## Create a new git tag (use VERSION=v1.0.0)
	@if [ -z "$(VERSION)" ]; then echo "Usage: make tag VERSION=v1.0.0"; exit 1; fi
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)

release-check: ## Check if ready for release
	@echo "Checking release readiness..."
	@git diff --exit-code || (echo "Working directory not clean"; exit 1)
	@git diff --cached --exit-code || (echo "Staging area not clean"; exit 1)
	make check
	@echo "✅ Ready for release"

# Docker targets
docker-build: ## Build Docker image
	docker build -t kolosys/proton:latest .

docker-run: ## Run proton in Docker
	docker run --rm -v $$(pwd):/workspace kolosys/proton:latest generate

# GitHub Actions testing
act: ## Test GitHub Actions locally (requires act)
	act -j generate

# Benchmarks
bench: ## Run benchmarks
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# Version info
version: ## Show version information
	@echo "Go version: $$(go version)"
	@echo "Git commit: $$(git rev-parse --short HEAD)"
	@echo "Git branch: $$(git rev-parse --abbrev-ref HEAD)"
	@if [ -f bin/proton ]; then echo "Proton version: $$(./bin/proton --version)"; fi
