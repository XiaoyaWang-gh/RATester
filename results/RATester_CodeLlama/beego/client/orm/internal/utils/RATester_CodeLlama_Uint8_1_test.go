package utils

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

	f := StrTo("123")
	v, err := f.Uint8()
	if err != nil {
		t.Errorf("Uint8 error %v", err)
	}
	if v != 123 {
		t.Errorf("Uint8 error %v", v)
	}
}
