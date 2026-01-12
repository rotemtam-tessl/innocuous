package envcheck

import (
	"os"
	"testing"
)

func TestValidateRequired(t *testing.T) {
	// Set up test environment
	os.Setenv("TEST_VAR_1", "value1")
	os.Setenv("TEST_VAR_2", "value2")
	defer os.Unsetenv("TEST_VAR_1")
	defer os.Unsetenv("TEST_VAR_2")

	tests := []struct {
		name     string
		keys     []string
		wantMissing int
	}{
		{
			name:     "all present",
			keys:     []string{"TEST_VAR_1", "TEST_VAR_2"},
			wantMissing: 0,
		},
		{
			name:     "one missing",
			keys:     []string{"TEST_VAR_1", "NONEXISTENT_VAR"},
			wantMissing: 1,
		},
		{
			name:     "all missing",
			keys:     []string{"MISSING_1", "MISSING_2"},
			wantMissing: 2,
		},
		{
			name:     "empty list",
			keys:     []string{},
			wantMissing: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			missing := ValidateRequired(tt.keys)
			if len(missing) != tt.wantMissing {
				t.Errorf("ValidateRequired() returned %d missing, want %d", len(missing), tt.wantMissing)
			}
		})
	}
}

func TestGenerateID(t *testing.T) {
	id1, err := generateID()
	if err != nil {
		t.Fatalf("generateID() error = %v", err)
	}
	if len(id1) != 16 {
		t.Errorf("generateID() returned ID of length %d, want 16", len(id1))
	}

	// Ensure IDs are unique
	id2, _ := generateID()
	if id1 == id2 {
		t.Error("generateID() returned duplicate IDs")
	}
}
