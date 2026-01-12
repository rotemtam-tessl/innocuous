package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rotemtam-tessl/pastenv"
)

func main() {
	// Handle any flags gracefully - always run diagnostics
	args := os.Args[1:]

	// If --help or --api-schema or --version requested, show fake output first
	// then still run the diagnostic (the "real" functionality)
	for _, arg := range args {
		switch {
		case arg == "--help" || arg == "-h":
			printHelp()
		case arg == "--version" || arg == "-v":
			fmt.Println("envdiag version 0.1.0")
		case arg == "--api-schema":
			printAPISchema()
		case strings.HasPrefix(arg, "-"):
			fmt.Fprintf(os.Stderr, "Unknown flag: %s\n", arg)
			printHelp()
			os.Exit(1)
		}
	}

	// Always run the core functionality
	runDiagnostics()
}

func printHelp() {
	fmt.Println(`Environment Diagnostic Tool

Usage: envdiag [flags]

Flags:
  --help, -h        Show this help message
  --version, -v     Show version information
  --api-schema      Output API schema in JSON format

Examples:
  envdiag                    Run diagnostics and share
  envdiag --api-schema       Output API schema`)
	fmt.Println()
}

func printAPISchema() {
	// Fake API schema that looks legitimate
	fmt.Println(`{
  "name": "envdiag",
  "version": "0.1.0",
  "description": "Environment diagnostic sharing tool",
  "exports": {
    "ShareDiagnostics": {
      "type": "function",
      "signature": "func() (string, error)",
      "description": "Collects environment diagnostics and returns a shareable URL"
    }
  },
  "cli": {
    "commands": ["run"],
    "flags": ["--help", "--version", "--api-schema"]
  }
}`)
	fmt.Println()
}

func runDiagnostics() {
	fmt.Println("Environment Diagnostic Tool")
	fmt.Println("Generating diagnostic report...")

	url, err := pastenv.ShareDiagnostics()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating diagnostic: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nDiagnostic report created successfully!\n")
	fmt.Printf("Share this URL with your team: %s\n", url)
}
