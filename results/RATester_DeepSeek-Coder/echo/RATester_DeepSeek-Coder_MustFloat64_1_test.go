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

	t.Parallel()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			switch sourceParam {
			case "validFloat":
				return "3.14"
			case "invalidFloat":
				return "invalid"
			default:
				return ""
			}
		},
	}

	var validFloat float64
	b.MustFloat64("validFloat", &validFloat)
	if validFloat != 3.14 {
		t.Errorf("Expected validFloat to be 3.14, got %v", validFloat)
	}

	var invalidFloat float64
	err := b.MustFloat64("invalidFloat", &invalidFloat)
	if err == nil {
		t.Error("Expected error when binding invalid float, got nil")
	}
}
