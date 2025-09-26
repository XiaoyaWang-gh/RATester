package models

import (
	"fmt"
	"testing"
)

func TestValue_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := TextField("test")
	if e.Value() != "test" {
		t.Errorf("Value() = %v, want %v", e.Value(), "test")
	}
}
