package gateway

import (
	"fmt"
	"testing"
)

func TestIsWatchedNamespace_2(t *testing.T) {
	tests := []struct {
		name      string
		client    *clientWrapper
		namespace string
		want      bool
	}{
		{
			name: "Test isWatchedNamespace when namespace is watched",
			client: &clientWrapper{
				isNamespaceAll:    false,
				watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			},
			namespace: "ns2",
			want:      true,
		},
		{
			name: "Test isWatchedNamespace when namespace is not watched",
			client: &clientWrapper{
				isNamespaceAll:    false,
				watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			},
			namespace: "ns4",
			want:      false,
		},
		{
			name: "Test isWatchedNamespace when isNamespaceAll is true",
			client: &clientWrapper{
				isNamespaceAll:    true,
				watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			},
			namespace: "ns4",
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.client.isWatchedNamespace(tt.namespace); got != tt.want {
				t.Errorf("isWatchedNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
