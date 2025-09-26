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

	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "ids" {
				return []string{"true", "false"}
			}
			return nil
		},
	}

	var dest []bool
	b.MustBools("ids", &dest)

	if len(dest) != 2 {
		t.Errorf("Expected length of dest to be 2, got %d", len(dest))
	}

	if !dest[0] {
		t.Errorf("Expected first element of dest to be true, got false")
	}

	if dest[1] {
		t.Errorf("Expected second element of dest to be false, got true")
	}
}
