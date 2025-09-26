package cache

import (
	"fmt"
	"testing"
)

func TestNew_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := New(10)
	if c.MaxEntries != 10 {
		t.Errorf("New() = %v, want %v", c.MaxEntries, 10)
	}
}
