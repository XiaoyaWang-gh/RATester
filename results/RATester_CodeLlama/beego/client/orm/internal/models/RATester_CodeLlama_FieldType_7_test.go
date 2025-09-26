package models

import (
	"fmt"
	"testing"
)

func TestFieldType_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &DateTimeField{}
	if e.FieldType() != TypeDateTimeField {
		t.Errorf("FieldType() = %v, want %v", e.FieldType(), TypeDateTimeField)
	}
}
