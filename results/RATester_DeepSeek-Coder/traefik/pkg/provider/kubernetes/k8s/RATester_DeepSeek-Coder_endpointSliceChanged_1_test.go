package k8s

import (
	"fmt"
	"testing"

	discoveryv1 "k8s.io/api/discovery/v1"
)

func TestEndpointSliceChanged_1(t *testing.T) {
	type args struct {
		a *discoveryv1.EndpointSlice
		b *discoveryv1.EndpointSlice
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := endpointSliceChanged(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("endpointSliceChanged() = %v, want %v", got, tt.want)
			}
		})
	}
}
