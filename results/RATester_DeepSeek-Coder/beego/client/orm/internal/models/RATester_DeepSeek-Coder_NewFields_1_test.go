package models

import (
	"fmt"
	"testing"
)

func TestNewFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewFields()

	if f == nil {
		t.Errorf("NewFields() = %v; want not nil", f)
	}

	if f.Fields == nil {
		t.Errorf("NewFields().Fields = %v; want not nil", f.Fields)
	}

	if f.FieldsLow == nil {
		t.Errorf("NewFields().FieldsLow = %v; want not nil", f.FieldsLow)
	}

	if f.Columns == nil {
		t.Errorf("NewFields().Columns = %v; want not nil", f.Columns)
	}

	if f.FieldsByType == nil {
		t.Errorf("NewFields().FieldsByType = %v; want not nil", f.FieldsByType)
	}
}
