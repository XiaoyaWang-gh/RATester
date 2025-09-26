package models

import (
	"fmt"
	"testing"
)

func TestValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := FloatField(1.23)
	if e.Value() != 1.23 {
		t.Errorf("e.Value() = %v, want %v", e.Value(), 1.23)
	}
}
