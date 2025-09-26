package utils

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGetGOPATHs_1(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		expected []string
	}{
		{
			name:     "Test with GOPATH set",
			env:      "/usr/local/go:/usr/local/bin",
			expected: []string{"/usr/local/go", "/usr/local/bin"},
		},
		{
			name:     "Test with empty GOPATH",
			env:      "",
			expected: []string{defaultGOPATH()},
		},
		{
			name:     "Test with no GOPATH",
			env:      "go1.8",
			expected: []string{defaultGOPATH()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv("GOPATH", tt.env)
			got := GetGOPATHs()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GetGOPATHs() = %v, want %v", got, tt.expected)
			}
		})
	}
}
