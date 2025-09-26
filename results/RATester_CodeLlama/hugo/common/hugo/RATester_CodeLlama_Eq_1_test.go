package hugo

import (
	"fmt"
	"testing"
)

func TestEq_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := VersionString("1.0.0")
	other := "1.0.0"
	if !h.Eq(other) {
		t.Errorf("h.Eq(other) = %v, want %v", false, true)
	}
}
