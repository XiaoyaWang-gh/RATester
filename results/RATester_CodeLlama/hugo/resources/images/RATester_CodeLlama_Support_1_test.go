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
		name:    "test",
		support: 1.0,
		kernel:  func(float32) float32 { return 1.0 },
	}
	if r.Support() != 1.0 {
		t.Errorf("Support() = %f, want %f", r.Support(), 1.0)
	}
}
