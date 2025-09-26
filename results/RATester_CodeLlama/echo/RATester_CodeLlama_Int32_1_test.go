package echo

import (
	"fmt"
	"testing"
)

func TestInt32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var dest int32
	b := new(ValueBinder)
	b.Int32("sourceParam", &dest)

	if len(b.errors) != 0 {
		t.Errorf("errors should be empty, got %v", b.errors)
	}

	if dest != 0 {
		t.Errorf("dest should be 0, got %v", dest)
	}
}
