BINARY_NAME=ctx
VERSION=0.1.0
MOD_NAME=github.com/harshalve/ctx-tool

.PHONY: build install clean test fmt vet precommit

# 1. Format the code according to Go standards
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 2. Run static analysis (checks for suspicious constructs)
vet:
	@echo "Vetting code..."
	go vet ./...

# 3. Run tests (if they exist, skips if none found)
test:
	@echo "Running tests..."
	go test ./...

# 4. Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	go build -o bin/$(BINARY_NAME) ./cmd/ctx/main.go

# 5. The Pre-Commit Command (The Gatekeeper)
precommit: fmt vet build
	@echo "✅ Pre-commit checks passed!"

# 6. Install globally
install: precommit
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# 7. Clean up
clean:
	@echo "Cleaning artifacts..."
	rm -rf bin/