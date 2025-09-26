package echo

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "sourceParam" {
				return "value"
			}
			return ""
		},
		failFast: true,
	}

	var dest string
	b.String("sourceParam", &dest)

	if dest != "value" {
		t.Errorf("Expected dest to be 'value', but got '%s'", dest)
	}

	b.String("otherParam", &dest)

	if dest != "value" {
		t.Errorf("Expected dest to be 'value', but got '%s'", dest)
	}
}
