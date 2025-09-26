package utils

import (
	"fmt"
	"testing"
)

func TestClear_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("1")
	f.Clear()
	if f.String() != string(rune(0x1E)) {
		t.Errorf("f.String() = %v, want %v", f.String(), string(rune(0x1E)))
	}
}
