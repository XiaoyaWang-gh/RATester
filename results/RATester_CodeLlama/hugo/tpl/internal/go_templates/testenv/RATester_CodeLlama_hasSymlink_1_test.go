package testenv

import (
	"fmt"
	"testing"
)

func TestHasSymlink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ok, reason := hasSymlink()
	if !ok {
		t.Errorf("hasSymlink() = (%v, %q), want (true, \"\")", ok, reason)
	}
}
