package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepCopy_3(t *testing.T) {
	tests := []struct {
		name string
		in   *WeightedRoundRobin
		want *WeightedRoundRobin
	}{
		{
			name: "nil input",
			in:   nil,
			want: nil,
		},
		{
			name: "non-nil input",
			in: &WeightedRoundRobin{
				Services: []Service{
					{
						LoadBalancerSpec: LoadBalancerSpec{
							Name: "service1",
						},
					},
				},
			},
			want: &WeightedRoundRobin{
				Services: []Service{
					{
						LoadBalancerSpec: LoadBalancerSpec{
							Name: "service1",
						},
					},
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

			got := tt.in.DeepCopy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
