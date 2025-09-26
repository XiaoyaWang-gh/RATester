package models

import (
	"fmt"
	"testing"
)

func TestValue_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := PositiveIntegerField(1)
	if got := e.Value(); got != 1 {
		t.Errorf("PositiveIntegerField.Value() = %v, want %v", got, 1)
	}
}
