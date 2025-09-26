package nomad

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildProviders_1(t *testing.T) {
	tests := []struct {
		name     string
		builder  *ProviderBuilder
		expected []*Provider
	}{
		{
			name: "no namespaces",
			builder: &ProviderBuilder{
				Namespaces: []string{},
			},
			expected: []*Provider{
				{
					Configuration: Configuration{},
					name:          providerName,
				},
			},
		},
		{
			name: "with namespaces",
			builder: &ProviderBuilder{
				Namespaces: []string{"ns1", "ns2"},
			},
			expected: []*Provider{
				{
					Configuration: Configuration{},
					name:          providerName + "-ns1",
					namespace:     "ns1",
				},
				{
					Configuration: Configuration{},
					name:          providerName + "-ns2",
					namespace:     "ns2",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			providers := test.builder.BuildProviders()
			if !reflect.DeepEqual(providers, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, providers)
			}
		})
	}
}
