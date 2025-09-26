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
	if e.FieldType() != TypeDateField {
		t.Errorf("FieldType() = %v, want %v", e.FieldType(), TypeDateField)
	}
}
