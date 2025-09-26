package models

import (
	"fmt"
	"testing"
)

func TestGetByAny_1(t *testing.T) {
	fields := &Fields{
		Fields: map[string]*FieldInfo{
			"field1": {Name: "field1"},
		},
		FieldsLow: map[string]*FieldInfo{
			"field1low": {Name: "field1low"},
		},
		Columns: map[string]*FieldInfo{
			"column1": {Name: "column1"},
		},
	}

	t.Run("ExistingField", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fi, ok := fields.GetByAny("field1")
		if !ok {
			t.Errorf("Expected to find field1")
		}
		if fi.Name != "field1" {
			t.Errorf("Expected field1, got %s", fi.Name)
		}
	})

	t.Run("ExistingFieldLowercase", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fi, ok := fields.GetByAny("Field1")
		if !ok {
			t.Errorf("Expected to find Field1")
		}
		if fi.Name != "field1low" {
			t.Errorf("Expected field1low, got %s", fi.Name)
		}
	})

	t.Run("ExistingColumn", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fi, ok := fields.GetByAny("column1")
		if !ok {
			t.Errorf("Expected to find column1")
		}
		if fi.Name != "column1" {
			t.Errorf("Expected column1, got %s", fi.Name)
		}
	})

	t.Run("NonExistingField", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, ok := fields.GetByAny("nonexisting")
		if ok {
			t.Errorf("Expected not to find nonexisting")
		}
	})
}
