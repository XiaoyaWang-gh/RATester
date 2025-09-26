package render

import (
	"fmt"
	"testing"
)

func TestOrdinal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &hookBase{}
	c.ordinal = 1
	if c.Ordinal() != 1 {
		t.Errorf("Ordinal() = %v, want %v", c.Ordinal(), 1)
	}
}
