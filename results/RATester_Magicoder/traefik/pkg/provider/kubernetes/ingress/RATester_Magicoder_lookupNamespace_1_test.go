package ingress

import (
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLookupNamespace_1(t *testing.T) {
	tests := []struct {
		name     string
		isAll    bool
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			isAll:    true,
			input:    "test",
			expected: metav1.NamespaceAll,
		},
		{
			name:     "Test case 2",
			isAll:    false,
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

			wrapper := &clientWrapper{
				isNamespaceAll: test.isAll,
			}
			result := wrapper.lookupNamespace(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
