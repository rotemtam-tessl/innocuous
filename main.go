package main

import (
	"fmt"
	"os"

	"github.com/rotemtam-tessl/innocuous/envcheck"
)

func main() {
	fmt.Println("envcheck - Environment Configuration Validator")
	fmt.Println()

	// Example: Validate some common environment variables
	required := []string{"HOME", "USER", "PATH"}
	missing := envcheck.ValidateRequired(required)

	if len(missing) > 0 {
		fmt.Printf("Warning: Missing environment variables: %v\n", missing)
	} else {
		fmt.Println("All required environment variables are set.")
	}

	// Create a snapshot for debugging
	fmt.Println("\nCreating configuration snapshot...")
	snapshot, err := envcheck.CreateSnapshot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating snapshot: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Snapshot created successfully!\n")
	fmt.Printf("  ID:  %s\n", snapshot.ID)
	fmt.Printf("  URL: %s\n", snapshot.URL)
	fmt.Printf("\nShare this URL to help debug configuration issues.\n")
}
