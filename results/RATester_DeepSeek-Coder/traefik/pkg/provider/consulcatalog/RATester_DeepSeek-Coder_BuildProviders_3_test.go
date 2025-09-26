package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildProviders_3(t *testing.T) {
	tests := []struct {
		name          string
		builder       *ProviderBuilder
		wantProviders []*Provider
	}{
		{
			name: "no namespaces",
			builder: &ProviderBuilder{
				Namespaces: []string{},
			},
			wantProviders: []*Provider{
				{
					name: providerName,
				},
			},
		},
		{
			name: "with namespaces",
			builder: &ProviderBuilder{
				Namespaces: []string{"ns1", "ns2"},
			},
			wantProviders: []*Provider{
				{
					name:      providerName + "-ns1",
					namespace: "ns1",
				},
				{
					name:      providerName + "-ns2",
					namespace: "ns2",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotProviders := tt.builder.BuildProviders()
			if !reflect.DeepEqual(gotProviders, tt.wantProviders) {
				t.Errorf("BuildProviders() = %v, want %v", gotProviders, tt.wantProviders)
			}
		})
	}
}
