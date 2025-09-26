package echo

import (
	"fmt"
	"testing"
)

func TestMustInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "123"
			}
			return ""
		},
	}

	var dest int
	b.MustInt("test", &dest)

	if dest != 123 {
		t.Errorf("Expected 123, got %d", dest)
	}
}
