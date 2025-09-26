package models

import (
	"fmt"
	"testing"
)

func TestFieldType_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &TimeField{}
	expected := TypeDateField
	actual := e.FieldType()
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
