package echo

import (
	"fmt"
	"testing"
)

func TestMustFloat64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "123.456"
		},
	}

	var dest float64
	err := b.MustFloat64("param", &dest)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if dest != 123.456 {
		t.Errorf("Expected dest to be 123.456, but got: %v", dest)
	}
}
