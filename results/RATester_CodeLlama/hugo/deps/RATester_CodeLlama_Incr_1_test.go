package deps

import (
	"fmt"
	"testing"
)

func TestIncr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{}
	b.counter = 1
	if b.Incr() != 2 {
		t.Errorf("Incr() = %d, want 2", b.Incr())
	}
}
