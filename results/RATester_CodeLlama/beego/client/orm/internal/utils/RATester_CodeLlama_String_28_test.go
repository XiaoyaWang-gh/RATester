package utils

import (
	"fmt"
	"testing"
)

func TestString_28(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("test")
	if f.String() != "test" {
		t.Errorf("String() = %v, want %v", f.String(), "test")
	}
}
