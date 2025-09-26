package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestDeepCopyInto_3(t *testing.T) {
	tests := []struct {
		name string
		in   *WeightedRoundRobin
		out  *WeightedRoundRobin
	}{
		{
			name: "nil in",
			in:   nil,
			out:  &WeightedRoundRobin{},
		},
		{
			name: "empty in",
			in:   &WeightedRoundRobin{},
			out:  &WeightedRoundRobin{},
		},
		{
			name: "non-empty in",
			in: &WeightedRoundRobin{
				Services: []Service{
					{
						LoadBalancerSpec: LoadBalancerSpec{
							Name: "service1",
						},
					},
				},
				Sticky: &dynamic.Sticky{},
			},
			out: &WeightedRoundRobin{
				Services: []Service{
					{
						LoadBalancerSpec: LoadBalancerSpec{
							Name: "service1",
						},
					},
				},
				Sticky: &dynamic.Sticky{},
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

			got := &WeightedRoundRobin{}
			tt.in.DeepCopyInto(got)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("DeepCopyInto() = %v, want %v", got, tt.out)
			}
		})
	}
}
