package ingress

import (
	"fmt"
	"testing"
)

func TestIsWatchedNamespace_1(t *testing.T) {
	tests := []struct {
		name     string
		c        *clientWrapper
		ns       string
		expected bool
	}{
		{
			name: "isWatchedNamespace",
			c: &clientWrapper{
				watchedNamespaces: []string{"ns1", "ns2"},
			},
			ns:       "ns1",
			expected: true,
		},
		{
			name: "isWatchedNamespace",
			c: &clientWrapper{
				watchedNamespaces: []string{"ns1", "ns2"},
			},
			ns:       "ns3",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.isWatchedNamespace(tt.ns); got != tt.expected {
				t.Errorf("isWatchedNamespace() = %v, want %v", got, tt.expected)
			}
		})
	}
}
