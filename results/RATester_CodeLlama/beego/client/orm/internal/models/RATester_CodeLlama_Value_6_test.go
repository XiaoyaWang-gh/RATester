package models

import (
	"fmt"
	"testing"
)

func TestValue_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := PositiveSmallIntegerField(1)
	if e.Value() != 1 {
		t.Errorf("PositiveSmallIntegerField.Value() = %v, want %v", e.Value(), 1)
	}
}
