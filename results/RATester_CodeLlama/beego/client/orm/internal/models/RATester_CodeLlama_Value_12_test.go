package models

import (
	"fmt"
	"testing"
)

func TestValue_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := BooleanField(true)
	if e.Value() != true {
		t.Errorf("Value() = %v, want %v", e.Value(), true)
	}
}
