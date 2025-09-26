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

	v := Version{Major: 1, Minor: 2, PatchLevel: 3}
	next := v.NextPatchLevel(4)

	if next.Major != 1 || next.Minor != 2 || next.PatchLevel != 4 {
		t.Errorf("Expected next version to be 1.2.4, but got %d.%d.%d", next.Major, next.Minor, next.PatchLevel)
	}
}
