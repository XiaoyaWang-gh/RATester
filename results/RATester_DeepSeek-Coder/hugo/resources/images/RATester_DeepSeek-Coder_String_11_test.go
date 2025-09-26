package images

import (
	"fmt"
	"testing"
)

func TestString_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := resamp{
		name:    "Test",
		support: 1.0,
		kernel:  func(x float32) float32 { return 0.0 },
	}

	expected := "Test"
	actual := r.String()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
