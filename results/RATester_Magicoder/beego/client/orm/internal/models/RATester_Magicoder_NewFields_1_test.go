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
		t.Error("NewFields() should not return nil")
	}

	if f.Fields == nil {
		t.Error("Fields should not be nil")
	}

	if f.FieldsLow == nil {
		t.Error("FieldsLow should not be nil")
	}

	if f.Columns == nil {
		t.Error("Columns should not be nil")
	}

	if f.FieldsByType == nil {
		t.Error("FieldsByType should not be nil")
	}
}
