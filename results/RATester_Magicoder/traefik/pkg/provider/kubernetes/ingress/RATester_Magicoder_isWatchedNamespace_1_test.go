package ingress

import (
	"fmt"
	"testing"
)

func TestIsWatchedNamespace_1(t *testing.T) {
	tests := []struct {
		name           string
		clientWrapper  *clientWrapper
		namespace      string
		expectedResult bool
	}{
		{
			name: "Test case 1: Namespace is watched",
			clientWrapper: &clientWrapper{
				isNamespaceAll: false,
				watchedNamespaces: []string{
					"namespace1",
					"namespace2",
				},
			},
			namespace:      "namespace1",
			expectedResult: true,
		},
		{
			name: "Test case 2: Namespace is not watched",
			clientWrapper: &clientWrapper{
				isNamespaceAll: false,
				watchedNamespaces: []string{
					"namespace1",
					"namespace2",
				},
			},
			namespace:      "namespace3",
			expectedResult: false,
		},
		{
			name: "Test case 3: All namespaces are watched",
			clientWrapper: &clientWrapper{
				isNamespaceAll: true,
			},
			namespace:      "namespace1",
			expectedResult: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.clientWrapper.isWatchedNamespace(test.namespace)
			if result != test.expectedResult {
				t.Errorf("Expected %v, but got %v", test.expectedResult, result)
			}
		})
	}
}
