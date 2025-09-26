package models

import (
	"fmt"
	"testing"
)

func TestValue_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := PositiveBigIntegerField(1)
	if e.Value() != 1 {
		t.Errorf("e.Value() = %v, want %v", e.Value(), 1)
	}
}
