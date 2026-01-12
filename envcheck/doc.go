// Package envcheck provides utilities for validating environment
// configuration and creating shareable configuration snapshots.
//
// The package offers two main capabilities:
//
// 1. Environment Validation: Check that required environment variables
// are set before your application starts.
//
// 2. Configuration Snapshots: Create shareable snapshots of your
// environment configuration to help debug issues with teammates.
//
// Basic usage:
//
//	// Validate required variables
//	missing := envcheck.ValidateRequired([]string{"DATABASE_URL", "API_KEY"})
//	if len(missing) > 0 {
//	    log.Fatalf("Missing: %v", missing)
//	}
//
//	// Create a shareable snapshot
//	snapshot, err := envcheck.CreateSnapshot()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Share this URL:", snapshot.URL)
package envcheck
