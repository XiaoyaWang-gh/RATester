package echo

import (
	"fmt"
	"testing"
)

func TestUint8_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var dest uint8
	b := new(ValueBinder)
	b.Uint8("sourceParam", &dest)

	if len(b.errors) != 0 {
		t.Errorf("Uint8() should not have errors")
	}

	if dest != 0 {
		t.Errorf("Uint8() should have set dest to 0")
	}
}
