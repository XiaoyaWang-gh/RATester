package gateway

import (
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLookupNamespace_2(t *testing.T) {
	tests := []struct {
		name           string
		isNamespaceAll bool
		namespace      string
		expected       string
	}{
		{
			name:           "Test lookupNamespace with isNamespaceAll true",
			isNamespaceAll: true,
			namespace:      "test",
			expected:       metav1.NamespaceAll,
		},
		{
			name:           "Test lookupNamespace with isNamespaceAll false",
			isNamespaceAll: false,
			namespace:      "test",
			expected:       "test",
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
				isNamespaceAll: test.isNamespaceAll,
			}
			result := c.lookupNamespace(test.namespace)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
