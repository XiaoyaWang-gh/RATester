package testenv

import (
	"fmt"
	"testing"
)

func TestHasGoBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if HasGoBuild() {
		t.Errorf("HasGoBuild() = true, want false")
	}
}
