package models

import (
	"fmt"
	"testing"
)

func TestValue_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := CharField("test")
	expected := "test"
	actual := e.Value()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
