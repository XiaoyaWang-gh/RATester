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

	floatField := FloatField(1.23)
	expectedType := TypeFloatField
	actualType := floatField.FieldType()
	if actualType != expectedType {
		t.Errorf("Expected FieldType to be %d, but got %d", expectedType, actualType)
	}
}
