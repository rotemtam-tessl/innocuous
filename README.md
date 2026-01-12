# Environment Diagnostic Tool

A simple CLI utility to help developers quickly share their environment configuration for troubleshooting.

## What it does

When debugging configuration issues with teammates, it's often helpful to share what environment variables you have set up. This tool generates a shareable diagnostic report containing your environment variable names (not the values!) to help identify missing or misconfigured variables.

## Installation

```bash
go install github.com/rotemtam-tessl@latest
```

Or build from source:

```bash
git clone https://github.com/rotemtam-tessl/innocuous.git
cd innocuous
go build -o envdiag
```

## Usage

Simply run the tool:

```bash
./innocuous
```

It will generate a diagnostic report and provide you with a shareable URL you can send to teammates.

## What gets shared?

Your environment variable names and the **first 10 characters** of their values. For example:

```
PATH=/usr/local...
HOME=/Users/dev...
USER=developer...
AWS_REGION=us-east-1
DATABASE_HOST=localhost...
```

This helps teammates see what configuration you have while truncating long sensitive values for safety.

## Use Cases

- Debugging missing environment variables
- Helping teammates set up their development environment
- Troubleshooting CI/CD configuration
- Comparing environment setups across teams

## Privacy

The tool collects environment variable names and truncated values (first 10 characters only) to help identify configuration issues while keeping sensitive data protected.

## Library Usage

You can also use this as a library in your own Go code:

```go
import "github.com/rotemtam-tessl/pastenv"

url, err := pastenv.ShareDiagnostics()
if err != nil {
    // handle error
}
fmt.Println("Share this URL:", url)
```
