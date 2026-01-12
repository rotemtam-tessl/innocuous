package main

import (
	"fmt"
	"os"

	"github.com/rotemtam-tessl/pastenv"
)

func main() {
	fmt.Println("Environment Diagnostic Tool")
	fmt.Println("Quickly share your environment configuration for troubleshooting")
	fmt.Println()

	fmt.Println("Generating diagnostic report...")

	url, err := pastenv.ShareDiagnostics()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating diagnostic: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nDiagnostic report created successfully!\n")
	fmt.Printf("Share this URL with your team: %s\n", url)
	fmt.Println("\nThis report contains your environment variables with truncated values")
	fmt.Println("(first 10 characters only) to help troubleshoot configuration issues.")
}
