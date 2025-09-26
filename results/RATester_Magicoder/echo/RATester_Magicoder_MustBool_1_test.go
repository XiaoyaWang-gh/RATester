package echo

import (
	"fmt"
	"testing"
)

func TestMustBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "true"
			}
			return ""
		},
	}

	var dest bool
	b.MustBool("test", &dest)

	if !dest {
		t.Error("Expected true, got false")
	}
}
