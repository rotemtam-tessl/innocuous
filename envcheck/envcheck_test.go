package envcheck

import (
	"os"
	"testing"
)

func TestValidateRequired(t *testing.T) {
	os.Setenv("TEST_VAR_1", "value1")
	os.Setenv("TEST_VAR_2", "value2")
	defer os.Unsetenv("TEST_VAR_1")
	defer os.Unsetenv("TEST_VAR_2")

	tests := []struct {
		name        string
		keys        []string
		wantMissing int
	}{
		{"all present", []string{"TEST_VAR_1", "TEST_VAR_2"}, 0},
		{"one missing", []string{"TEST_VAR_1", "NONEXISTENT"}, 1},
		{"all missing", []string{"MISSING_1", "MISSING_2"}, 2},
		{"empty list", []string{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			missing := ValidateRequired(tt.keys)
			if len(missing) != tt.wantMissing {
				t.Errorf("got %d missing, want %d", len(missing), tt.wantMissing)
			}
		})
	}
}

func TestValidateNonEmpty(t *testing.T) {
	os.Setenv("NON_EMPTY", "value")
	os.Setenv("WHITESPACE_ONLY", "   ")
	defer os.Unsetenv("NON_EMPTY")
	defer os.Unsetenv("WHITESPACE_ONLY")

	tests := []struct {
		name      string
		keys      []string
		wantEmpty int
	}{
		{"non-empty value", []string{"NON_EMPTY"}, 0},
		{"whitespace only", []string{"WHITESPACE_ONLY"}, 1},
		{"not set", []string{"NOT_SET_VAR"}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			empty := ValidateNonEmpty(tt.keys)
			if len(empty) != tt.wantEmpty {
				t.Errorf("got %d empty, want %d", len(empty), tt.wantEmpty)
			}
		})
	}
}

func TestGetEnvWithDefault(t *testing.T) {
	os.Setenv("EXISTING_VAR", "existing_value")
	defer os.Unsetenv("EXISTING_VAR")

	tests := []struct {
		name         string
		key          string
		defaultValue string
		want         string
	}{
		{"existing var", "EXISTING_VAR", "default", "existing_value"},
		{"missing var", "MISSING_VAR", "default", "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetEnvWithDefault(tt.key, tt.defaultValue)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
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
		t.Errorf("ID length = %d, want 16", len(id1))
	}

	id2, _ := generateID()
	if id1 == id2 {
		t.Error("generateID() returned duplicate IDs")
	}
}
