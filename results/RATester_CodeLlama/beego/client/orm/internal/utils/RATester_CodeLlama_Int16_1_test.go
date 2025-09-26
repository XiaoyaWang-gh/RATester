package utils

import (
	"fmt"
	"testing"
)

func TestInt16_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("1234")
	v, err := f.Int16()
	if err != nil {
		t.Errorf("Int16() error = %v", err)
		return
	}
	if v != 1234 {
		t.Errorf("Int16() = %v, want %v", v, 1234)
	}
}
