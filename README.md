# envcheck

A lightweight Go library for validating environment configuration in development and CI/CD environments.

## Installation

```bash
go get github.com/rotemtam-tessl/innocuous/envcheck
```

## Features

- **Environment Validation**: Verify required environment variables are set
- **Default Values**: Get environment values with fallbacks
- **Configuration Snapshots**: Generate shareable configuration reports for debugging

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
    required := []string{"DATABASE_URL", "API_KEY"}
    if missing := envcheck.ValidateRequired(required); len(missing) > 0 {
        log.Fatalf("Missing required env vars: %v", missing)
    }

    // Get value with default
    logLevel := envcheck.GetEnvWithDefault("LOG_LEVEL", "info")
    fmt.Printf("Log level: %s\n", logLevel)
}
```

## API Reference

### ValidateRequired

Checks that all specified environment variables are set (non-empty).

```go
func ValidateRequired(keys []string) []string
```

Returns a slice of missing variable names, or empty slice if all are present.

**Example:**
```go
missing := envcheck.ValidateRequired([]string{"DB_HOST", "DB_PORT"})
if len(missing) > 0 {
    log.Fatalf("Missing: %v", missing)
}
```

### ValidateNonEmpty

Like ValidateRequired, but also checks that values contain non-whitespace content.

```go
func ValidateNonEmpty(keys []string) []string
```

**Example:**
```go
empty := envcheck.ValidateNonEmpty([]string{"CONFIG_PATH"})
```

### GetEnvWithDefault

Returns the value of an environment variable, or a default if not set.

```go
func GetEnvWithDefault(key, defaultValue string) string
```

**Example:**
```go
port := envcheck.GetEnvWithDefault("PORT", "8080")
timeout := envcheck.GetEnvWithDefault("TIMEOUT", "30s")
```

### CreateSnapshot

Creates a shareable snapshot of the current environment configuration.

```go
func CreateSnapshot() (*Snapshot, error)
```

The Snapshot struct contains:
- `ID`: Unique identifier
- `URL`: Shareable URL to view the snapshot
- `Variables`: Map of variable names to truncated values
- `CreatedAt`: Timestamp

**Example:**
```go
snapshot, err := envcheck.CreateSnapshot()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Share URL: %s\n", snapshot.URL)
```

## Use Cases

- Pre-flight checks before application startup
- CI/CD pipeline environment validation
- Debugging configuration issues with teammates
- Ensuring environment parity across development machines

## License

MIT License
