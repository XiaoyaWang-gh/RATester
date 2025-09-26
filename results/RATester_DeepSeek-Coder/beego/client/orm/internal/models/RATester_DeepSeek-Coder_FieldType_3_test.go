package models

import (
	"fmt"
	"testing"
)

func TestFieldType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dateField := &DateField{}
	expected := TypeDateField
	actual := dateField.FieldType()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
