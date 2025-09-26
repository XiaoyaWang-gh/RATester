package utils

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

	f := StrTo("12345")
	v, err := f.Uint16()
	if err != nil {
		t.Errorf("Uint16 error %v", err)
	}
	if v != 12345 {
		t.Errorf("Uint16 error %v", v)
	}
}
