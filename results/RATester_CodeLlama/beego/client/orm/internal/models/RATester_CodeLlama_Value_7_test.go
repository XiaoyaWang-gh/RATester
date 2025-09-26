package models

import (
	"fmt"
	"testing"
)

func TestValue_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := SmallIntegerField(-32768)
	if e.Value() != -32768 {
		t.Errorf("Value() = %v, want %v", e.Value(), -32768)
	}
}
