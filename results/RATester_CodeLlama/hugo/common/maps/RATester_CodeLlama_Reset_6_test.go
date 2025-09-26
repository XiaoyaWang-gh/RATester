package maps

import (
	"fmt"
	"testing"
)

func TestReset_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &SliceCache[string]{}
	c.Append("key", "value")
	c.Reset()
	if len(c.m) != 0 {
		t.Errorf("expected empty map, got %v", c.m)
	}
}
