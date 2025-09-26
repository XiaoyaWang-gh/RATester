package utils

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

	f := StrTo("123")
	v, err := f.Int32()
	if err != nil {
		t.Errorf("Int32() error = %v", err)
		return
	}
	if v != 123 {
		t.Errorf("Int32() = %v, want %v", v, 123)
	}
}
