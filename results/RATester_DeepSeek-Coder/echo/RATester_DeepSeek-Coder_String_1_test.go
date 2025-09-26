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
			if sourceParam == "test" {
				return "test_value"
			}
			return ""
		},
	}

	var dest string
	b.String("test", &dest)
	if dest != "test_value" {
		t.Errorf("Expected dest to be 'test_value', got '%s'", dest)
	}

	b.String("not_test", &dest)
	if dest != "" {
		t.Errorf("Expected dest to be '', got '%s'", dest)
	}
}
