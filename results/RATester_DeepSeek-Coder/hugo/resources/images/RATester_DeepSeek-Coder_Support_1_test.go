package images

import (
	"fmt"
	"testing"
)

func TestSupport_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := resamp{
		name:    "TestResamp",
		support: 1.5,
		kernel:  func(x float32) float32 { return x },
	}

	support := r.Support()

	if support != r.support {
		t.Errorf("Expected support to be %v, but got %v", r.support, support)
	}
}
