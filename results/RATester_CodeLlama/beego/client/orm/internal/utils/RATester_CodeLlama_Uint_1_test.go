package utils

import (
	"fmt"
	"testing"
)

func TestUint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("123")
	v, err := f.Uint()
	if err != nil {
		t.Errorf("Uint() error = %v", err)
		return
	}
	if v != 123 {
		t.Errorf("Uint() = %v, want %v", v, 123)
	}
}
