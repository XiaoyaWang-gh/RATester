package echo

import (
	"fmt"
	"testing"
)

func TestMustFloat32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "3.14"
		},
	}

	var f float32
	err := b.MustFloat32("pi", &f).BindError()

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if f != float32(3.14) {
		t.Errorf("Expected f to be 3.14, but got: %v", f)
	}
}
