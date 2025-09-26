package crd

import (
	"fmt"
	"testing"
)

func TestIsWatchedNamespace_3(t *testing.T) {
	tests := []struct {
		name              string
		watchedNamespaces []string
		isNamespaceAll    bool
		ns                string
		want              bool
	}{
		{
			name:              "namespace is watched",
			watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			isNamespaceAll:    false,
			ns:                "ns2",
			want:              true,
		},
		{
			name:              "namespace is not watched",
			watchedNamespaces: []string{"ns1", "ns2", "ns3"},
			isNamespaceAll:    false,
			ns:                "ns4",
			want:              false,
		},
		{
			name:              "all namespaces are watched",
			watchedNamespaces: []string{},
			isNamespaceAll:    true,
			ns:                "ns4",
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
			if got := c.isWatchedNamespace(tt.ns); got != tt.want {
				t.Errorf("isWatchedNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
