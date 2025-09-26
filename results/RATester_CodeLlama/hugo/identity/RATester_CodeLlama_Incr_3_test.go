package identity

import (
	"fmt"
	"testing"
)

func TestIncr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IncrementByOne{}
	if got := c.Incr(); got != 1 {
		t.Errorf("IncrementByOne.Incr() = %v, want 1", got)
	}
}
