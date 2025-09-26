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
		support: 0.5,
		kernel:  func(x float32) float32 { return x },
	}

	expected := float32(0.5)
	result := r.Support()

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}
