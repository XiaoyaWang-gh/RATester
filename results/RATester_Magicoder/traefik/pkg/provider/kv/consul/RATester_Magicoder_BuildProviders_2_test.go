package consul

import (
	"fmt"
	"testing"
)

func TestBuildProviders_2(t *testing.T) {
	tests := []struct {
		name        string
		builder     *ProviderBuilder
		expectedLen int
	}{
		{
			name: "no namespaces",
			builder: &ProviderBuilder{
				Namespaces: []string{},
			},
			expectedLen: 1,
		},
		{
			name: "with namespaces",
			builder: &ProviderBuilder{
				Namespaces: []string{"ns1", "ns2"},
			},
			expectedLen: 2,
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
			if len(providers) != test.expectedLen {
				t.Errorf("Expected %d providers, got %d", test.expectedLen, len(providers))
			}
		})
	}
}
