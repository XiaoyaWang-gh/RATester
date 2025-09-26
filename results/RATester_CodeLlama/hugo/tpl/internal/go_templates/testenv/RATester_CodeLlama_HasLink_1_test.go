package testenv

import (
	"fmt"
	"runtime"
	"testing"
)

func TestHasLink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if runtime.GOOS == "plan9" || runtime.GOOS == "android" {
		if HasLink() {
			t.Errorf("HasLink() = true, want false")
		}
	} else {
		if !HasLink() {
			t.Errorf("HasLink() = false, want true")
		}
	}
}
