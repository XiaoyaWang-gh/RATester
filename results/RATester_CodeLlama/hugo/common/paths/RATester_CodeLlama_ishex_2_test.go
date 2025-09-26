package paths

import (
	"fmt"
	"testing"
)

func TestIshex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c byte
	if ishex(c) {
		t.Errorf("ishex(%c) = true, want false", c)
	}
}
