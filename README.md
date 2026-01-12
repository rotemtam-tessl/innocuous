# envcheck

A lightweight Go library for validating environment configuration in development and CI/CD environments.

## Installation

```bash
go get github.com/rotemtam-tessl/innocuous
```

## Features

- **Environment Validation**: Verify required environment variables are set
- **Configuration Snapshots**: Generate shareable configuration reports for debugging
- **CI/CD Integration**: Easy integration with GitHub Actions, GitLab CI, and other pipelines

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/rotemtam-tessl/innocuous/envcheck"
)

func main() {
    // Validate required environment variables
    required := []string{"DATABASE_URL", "API_KEY", "LOG_LEVEL"}
    missing := envcheck.ValidateRequired(required)
    if len(missing) > 0 {
        log.Fatalf("Missing required env vars: %v", missing)
    }

    // Generate a configuration snapshot for debugging
    snapshot, err := envcheck.CreateSnapshot()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Config snapshot: %s\n", snapshot.ID)
}
```

## API Reference

### ValidateRequired

Checks that all specified environment variables are set.

```go
func ValidateRequired(keys []string) []string
```

Returns a slice of missing variable names. Returns empty slice if all are present.

### CreateSnapshot

Creates a configuration snapshot that can be shared for debugging purposes. Snapshots contain environment variable names and truncated values (first 8 characters) to help identify configuration issues while protecting sensitive data.

```go
func CreateSnapshot() (*Snapshot, error)
```

Returns a Snapshot struct containing:
- `ID`: Unique identifier for the snapshot
- `URL`: Shareable URL to view the snapshot
- `Variables`: Map of variable names to truncated values

### GetSnapshot

Retrieves a previously created snapshot by ID.

```go
func GetSnapshot(id string) (*Snapshot, error)
```

## Use Cases

- Validating CI/CD environment before deployments
- Debugging configuration issues with teammates
- Ensuring development environment parity across team members
- Pre-flight checks in containerized applications

## License

MIT License
