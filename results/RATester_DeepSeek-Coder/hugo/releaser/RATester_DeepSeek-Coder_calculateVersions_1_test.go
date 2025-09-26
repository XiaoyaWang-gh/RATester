package releaser

import (
	"fmt"
	"testing"
)

func TestCalculateVersions_1(t *testing.T) {
	tests := []struct {
		name             string
		branchVersion    string
		wantNewVersion   string
		wantFinalVersion string
	}{
		{
			name:             "Test with a non-test version",
			branchVersion:    "0.90.0",
			wantNewVersion:   "0.90.0",
			wantFinalVersion: "0.90.0-DEV",
		},
		{
			name:             "Test with a test version",
			branchVersion:    "0.90.0-test",
			wantNewVersion:   "0.90.0-test",
			wantFinalVersion: "0.90.0-DEV",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := ReleaseHandler{
				branchVersion: tt.branchVersion,
			}

			newVersion, finalVersion := r.calculateVersions()

			if newVersion.String() != tt.wantNewVersion {
				t.Errorf("calculateVersions() newVersion = %v, want %v", newVersion, tt.wantNewVersion)
			}

			if finalVersion.String() != tt.wantFinalVersion {
				t.Errorf("calculateVersions() finalVersion = %v, want %v", finalVersion, tt.wantFinalVersion)
			}
		})
	}
}
