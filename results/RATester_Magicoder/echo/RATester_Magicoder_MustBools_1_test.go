package echo

import (
	"fmt"
	"testing"
)

func TestMustBools_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	binder := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "bools" {
				return []string{"true", "false"}
			}
			return nil
		},
	}

	var bools []bool
	binder.MustBools("bools", &bools)

	if len(bools) != 2 {
		t.Errorf("Expected 2 bools, got %d", len(bools))
	}
	if !bools[0] {
		t.Errorf("Expected first bool to be true, got false")
	}
	if bools[1] {
		t.Errorf("Expected second bool to be false, got true")
	}
}
