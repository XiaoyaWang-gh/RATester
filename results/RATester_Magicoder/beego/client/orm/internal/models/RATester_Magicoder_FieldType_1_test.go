package models

import (
	"fmt"
	"testing"
)

func TestFieldType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	field := BooleanField(true)
	expected := TypeBooleanField
	actual := field.FieldType()
	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
