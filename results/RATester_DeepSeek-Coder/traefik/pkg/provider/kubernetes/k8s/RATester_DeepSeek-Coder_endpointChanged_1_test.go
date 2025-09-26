package k8s

import (
	"fmt"
	"testing"

	discoveryv1 "k8s.io/api/discovery/v1"
)

func TestEndpointChanged_1(t *testing.T) {
	type args struct {
		a discoveryv1.Endpoint
		b discoveryv1.Endpoint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "addresses are different",
			args: args{
				a: discoveryv1.Endpoint{
					Addresses: []string{"192.0.2.1"},
				},
				b: discoveryv1.Endpoint{
					Addresses: []string{"192.0.2.2"},
				},
			},
			want: true,
		},
		{
			name: "addresses are the same",
			args: args{
				a: discoveryv1.Endpoint{
					Addresses: []string{"192.0.2.1"},
				},
				b: discoveryv1.Endpoint{
					Addresses: []string{"192.0.2.1"},
				},
			},
			want: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := endpointChanged(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("endpointChanged() = %v, want %v", got, tt.want)
			}
		})
	}
}
