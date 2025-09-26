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
	expected := true
	actual := e.Value()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
