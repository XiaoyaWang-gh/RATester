package models

import (
	"fmt"
	"testing"
)

func TestFieldType_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	floatField := new(FloatField)
	expected := TypeFloatField
	actual := floatField.FieldType()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
