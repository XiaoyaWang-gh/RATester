package models

import (
	"fmt"
	"testing"
)

func TestNewFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewFields()
	if f == nil {
		t.Error("NewFields() failed")
	}
}
