package utils

import (
	"fmt"
	"testing"
)

func TestFloat32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("1.23456789")
	v, err := f.Float32()
	if err != nil {
		t.Errorf("Float32() error = %v", err)
		return
	}
	if v != 1.23456789 {
		t.Errorf("Float32() = %v, want %v", v, 1.23456789)
	}
}
