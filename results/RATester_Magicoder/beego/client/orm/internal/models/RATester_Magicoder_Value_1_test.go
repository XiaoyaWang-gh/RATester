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

	e := FloatField(123.456)
	expected := 123.456
	actual := e.Value()
	if actual != expected {
		t.Errorf("Expected %f, but got %f", expected, actual)
	}
}
