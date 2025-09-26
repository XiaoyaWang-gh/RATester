package ingress

import (
	"fmt"
	"testing"
)

func TestIsWatchedNamespace_1(t *testing.T) {
	tests := []struct {
		name              string
		watchedNamespaces []string
		isNamespaceAll    bool
		input             string
		want              bool
	}{
		{
			name:              "Test case 1",
			watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			isNamespaceAll:    false,
			input:             "ns1",
			want:              true,
		},
		{
			name:              "Test case 2",
			watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			isNamespaceAll:    false,
			input:             "ns4",
			want:              false,
		},
		{
			name:              "Test case 3",
			watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			isNamespaceAll:    true,
			input:             "ns4",
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
			got := c.isWatchedNamespace(tt.input)
			if got != tt.want {
				t.Errorf("isWatchedNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
