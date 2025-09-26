package hugo

import (
	"fmt"
	"testing"
)

func TestNextPatchLevel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := Version{
		Major:      1,
		Minor:      2,
		PatchLevel: 3,
	}

	next := v.NextPatchLevel(4)

	if next.PatchLevel != 4 {
		t.Errorf("Expected patch level 4, got %d", next.PatchLevel)
	}

	if next.Major != 1 {
		t.Errorf("Expected major 1, got %d", next.Major)
	}

	if next.Minor != 2 {
		t.Errorf("Expected minor 2, got %d", next.Minor)
	}
}
