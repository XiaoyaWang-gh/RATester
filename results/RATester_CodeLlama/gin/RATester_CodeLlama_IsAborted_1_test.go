package gin

import (
	"fmt"
	"testing"
)

func TestIsAborted_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{index: 1}
	if !c.IsAborted() {
		t.Errorf("IsAborted() = %v, want %v", c.IsAborted(), true)
	}
}
