package utils

import (
	"fmt"
	"testing"
)

func TestFloat64_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("1.23456789")
	v, err := f.Float64()
	if err != nil {
		t.Errorf("Float64() error = %v", err)
		return
	}
	if v != 1.23456789 {
		t.Errorf("Float64() = %v, want %v", v, 1.23456789)
	}
}
