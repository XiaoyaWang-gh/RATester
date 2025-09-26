package v1alpha1

import (
	"fmt"
	"testing"
)

func TestDeepCopy_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := &WeightedRoundRobin{
		Services: []Service{
			{
				LoadBalancerSpec: LoadBalancerSpec{
					Name: "foo",
				},
			},
		},
	}
	out := in.DeepCopy()
	if out == in {
		t.Fatal("DeepCopy() should return a new object")
	}
	if out.Services[0].Name != "foo" {
		t.Errorf("DeepCopy() should have copied the name field")
	}
}
