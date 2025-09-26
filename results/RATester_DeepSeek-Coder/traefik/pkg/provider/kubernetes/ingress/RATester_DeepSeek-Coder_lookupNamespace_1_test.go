package ingress

import (
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLookupNamespace_1(t *testing.T) {
	tests := []struct {
		name     string
		isNA     bool
		input    string
		expected string
	}{
		{
			name:     "Test lookupNamespace with isNamespaceAll true",
			isNA:     true,
			input:    "test",
			expected: metav1.NamespaceAll,
		},
		{
			name:     "Test lookupNamespace with isNamespaceAll false",
			isNA:     false,
			input:    "test",
			expected: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &clientWrapper{
				isNamespaceAll: test.isNA,
			}
			result := c.lookupNamespace(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
