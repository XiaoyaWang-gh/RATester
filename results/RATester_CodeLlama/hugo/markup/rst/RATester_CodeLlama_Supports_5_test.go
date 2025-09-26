package rst

import (
	"fmt"
	"testing"
)

func TestSupports_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	// when
	// then
	if Supports() {
		t.Error("rst not installed")
	}
}
