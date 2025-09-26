package echo

import (
	"fmt"
	"testing"
)

func TestUint16_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var dest uint16
	b := new(ValueBinder)
	b.Uint16("sourceParam", &dest)

	if len(b.errors) != 0 {
		t.Errorf("Uint16() should not have errors")
	}

	if dest != 0 {
		t.Errorf("Uint16() should have set dest to 0")
	}
}
