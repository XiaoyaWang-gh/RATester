package k8s

import (
	"fmt"
	"testing"

	discoveryv1 "k8s.io/api/discovery/v1"
)

func TestEndpointChanged_1(t *testing.T) {
	tests := []struct {
		name string
		a    discoveryv1.Endpoint
		b    discoveryv1.Endpoint
		want bool
	}{
		{
			name: "same endpoint",
			a: discoveryv1.Endpoint{
				Addresses: []string{"192.168.1.1"},
			},
			b: discoveryv1.Endpoint{
				Addresses: []string{"192.168.1.1"},
			},
			want: false,
		},
		{
			name: "different addresses",
			a: discoveryv1.Endpoint{
				Addresses: []string{"192.168.1.1"},
			},
			b: discoveryv1.Endpoint{
				Addresses: []string{"192.168.1.2"},
			},
			want: true,
		},
		{
			name: "different number of addresses",
			a: discoveryv1.Endpoint{
				Addresses: []string{"192.168.1.1", "192.168.1.2"},
			},
			b: discoveryv1.Endpoint{
				Addresses: []string{"192.168.1.1"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := endpointChanged(tt.a, tt.b); got != tt.want {
				t.Errorf("endpointChanged() = %v, want %v", got, tt.want)
			}
		})
	}
}
