package common

import (
	"fmt"
	"testing"
)

func TestGetInstallPath_1(t *testing.T) {
	type test struct {
		name     string
		expected string
	}

	tests := []test{
		{
			name:     "windows",
			expected: `C:\Program Files\nps`,
		},
		{
			name:     "linux",
			expected: "/etc/nps",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetInstallPath(); got != tt.expected {
				t.Errorf("GetInstallPath() = %v, want %v", got, tt.expected)
			}
		})
	}
}
