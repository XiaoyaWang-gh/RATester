package testenv

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestHasCGO_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	origEnv = os.Environ()
	defer func() {
		os.Clearenv()
		for _, kv := range origEnv {
			parts := strings.SplitN(kv, "=", 2)
			os.Setenv(parts[0], parts[1])
		}
	}()

	tests := []struct {
		name     string
		env      map[string]string
		expected bool
	}{
		{
			name: "CGO_ENABLED is set to 1",
			env: map[string]string{
				"CGO_ENABLED": "1",
			},
			expected: true,
		},
		{
			name: "CGO_ENABLED is set to 0",
			env: map[string]string{
				"CGO_ENABLED": "0",
			},
			expected: false,
		},
		{
			name:     "CGO_ENABLED is not set",
			env:      map[string]string{},
			expected: false,
		},
	}

	for _, test := range tests {
		for k, v := range test.env {
			os.Setenv(k, v)
		}
		result := HasCGO()
		if result != test.expected {
			t.Errorf("Test case '%s' failed. Expected: %v, Got: %v", test.name, test.expected, result)
		}
	}
}
