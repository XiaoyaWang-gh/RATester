package echo

import (
	"fmt"
	"testing"
)

func TestMustFloat64s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1.23"
		},
	}

	var dest []float64
	err := b.MustFloat64s("param", &dest)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(dest) != 1 {
		t.Errorf("Expected dest to have 1 element, but got: %d", len(dest))
	}

	if dest[0] != 1.23 {
		t.Errorf("Expected dest[0] to be 1.23, but got: %f", dest[0])
	}
}
