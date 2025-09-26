package testenv

import (
	"fmt"
	"testing"
)

func TestHasGoRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if !HasGoRun() {
		t.Errorf("HasGoRun() = false, want true")
	}
}
