package gateway

import (
	"fmt"
	"testing"
)

func TestIsWatchedNamespace_2(t *testing.T) {
	tests := []struct {
		name              string
		watchedNamespaces []string
		isNamespaceAll    bool
		namespace         string
		want              bool
	}{
		{
			name:              "watched namespace",
			watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			isNamespaceAll:    false,
			namespace:         "ns2",
			want:              true,
		},
		{
			name:              "not watched namespace",
			watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			isNamespaceAll:    false,
			namespace:         "ns4",
			want:              false,
		},
		{
			name:              "all namespaces",
			watchedNamespaces: []string{},
			isNamespaceAll:    true,
			namespace:         "ns4",
			want:              true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &clientWrapper{
				watchedNamespaces: tt.watchedNamespaces,
				isNamespaceAll:    tt.isNamespaceAll,
			}
			if got := c.isWatchedNamespace(tt.namespace); got != tt.want {
				t.Errorf("isWatchedNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
