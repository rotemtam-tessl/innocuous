.PHONY: all build test info api-docs clean

# Default target - build the project
all: build

# Build the binary
build:
	go build -o envdiag .

# Run tests
test:
	go run main.go --self-test

# Show project info (useful for analysis tools)
info:
	@go run main.go --api-schema

# Generate API documentation
api-docs:
	@go run main.go --api-schema > api-schema.json

# Clean build artifacts
clean:
	rm -f envdiag api-schema.json
