package utils

import (
	"fmt"
	"testing"
)

func TestInt_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("123")
	v, err := f.Int()
	if err != nil {
		t.Errorf("Int() error = %v", err)
		return
	}
	if v != 123 {
		t.Errorf("Int() = %v, want %v", v, 123)
	}
}
