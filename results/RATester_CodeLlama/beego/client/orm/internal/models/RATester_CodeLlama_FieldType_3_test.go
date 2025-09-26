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

	e := &DateField{}
	if e.FieldType() != TypeDateField {
		t.Errorf("FieldType() = %v, want %v", e.FieldType(), TypeDateField)
	}
}
