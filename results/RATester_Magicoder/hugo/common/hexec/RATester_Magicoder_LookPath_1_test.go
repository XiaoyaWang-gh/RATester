package hexec

import (
	"fmt"
	"testing"
)

func TestLookPath_1(t *testing.T) {
	tests := []struct {
		name         string
		binaryName   string
		expectedPath string
	}{
		{
			name:         "valid binary",
			binaryName:   "ls",
			expectedPath: "/bin/ls",
		},
		{
			name:         "invalid binary",
			binaryName:   "invalid_binary",
			expectedPath: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			path := LookPath(tt.binaryName)
			if path != tt.expectedPath {
				t.Errorf("LookPath(%s) = %s; want %s", tt.binaryName, path, tt.expectedPath)
			}
		})
	}
}
