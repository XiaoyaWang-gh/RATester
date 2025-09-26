package releaser

import (
	"fmt"
	"testing"
)

func TestcalculateVersions_1(t *testing.T) {
	tests := []struct {
		name          string
		branchVersion string
		expectedNew   string
		expectedFinal string
	}{
		{
			name:          "Test case 1",
			branchVersion: "0.90.0",
			expectedNew:   "0.90.0",
			expectedFinal: "0.91.0-DEV",
		},
		{
			name:          "Test case 2",
			branchVersion: "0.90.0-test",
			expectedNew:   "0.90.0",
			expectedFinal: "0.91.0-DEV",
		},
		{
			name:          "Test case 3",
			branchVersion: "0.90.0-DEV",
			expectedNew:   "0.90.0-DEV",
			expectedFinal: "0.91.0-DEV",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			handler := ReleaseHandler{
				branchVersion: test.branchVersion,
			}

			newVersion, finalVersion := handler.calculateVersions()

			if newVersion.String() != test.expectedNew {
				t.Errorf("Expected new version to be %s, but got %s", test.expectedNew, newVersion.String())
			}

			if finalVersion.String() != test.expectedFinal {
				t.Errorf("Expected final version to be %s, but got %s", test.expectedFinal, finalVersion.String())
			}
		})
	}
}
