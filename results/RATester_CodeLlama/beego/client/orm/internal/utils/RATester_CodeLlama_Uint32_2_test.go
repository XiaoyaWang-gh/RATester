package utils

import (
	"fmt"
	"testing"
)

func TestUint32_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("1234567890")
	v, err := f.Uint32()
	if err != nil {
		t.Errorf("Uint32() error = %v", err)
		return
	}
	if v != 1234567890 {
		t.Errorf("Uint32() = %v, want %v", v, 1234567890)
	}
}
