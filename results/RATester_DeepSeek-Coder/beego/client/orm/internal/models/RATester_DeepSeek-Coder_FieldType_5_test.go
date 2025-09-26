package models

import (
	"fmt"
	"testing"
)

func TestFieldType_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := new(IntegerField)
	expected := TypeIntegerField
	actual := e.FieldType()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
