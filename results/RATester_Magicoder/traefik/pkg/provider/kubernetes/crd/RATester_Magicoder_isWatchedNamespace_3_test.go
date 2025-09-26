package crd

import (
	"fmt"
	"testing"
)

func TestIsWatchedNamespace_3(t *testing.T) {
	tests := []struct {
		name           string
		clientWrapper  clientWrapper
		namespace      string
		expectedResult bool
	}{
		{
			name: "Test isWatchedNamespace when namespace is watched",
			clientWrapper: clientWrapper{
				isNamespaceAll:    false,
				watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			},
			namespace:      "ns2",
			expectedResult: true,
		},
		{
			name: "Test isWatchedNamespace when namespace is not watched",
			clientWrapper: clientWrapper{
				isNamespaceAll:    false,
				watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			},
			namespace:      "ns4",
			expectedResult: false,
		},
		{
			name: "Test isWatchedNamespace when isNamespaceAll is true",
			clientWrapper: clientWrapper{
				isNamespaceAll:    true,
				watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			},
			namespace:      "ns4",
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
