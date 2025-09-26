package models

import (
	"fmt"
	"testing"
)

func TestGetByName_1(t *testing.T) {
	fields := &Fields{
		Fields: map[string]*FieldInfo{
			"field1": {Name: "field1"},
			"field2": {Name: "field2"},
		},
	}

	t.Run("existing field", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		field := fields.GetByName("field1")
		if field == nil {
			t.Fatal("expected field, got nil")
		}
		if field.Name != "field1" {
			t.Errorf("expected field1, got %s", field.Name)
		}
	})

	t.Run("non-existing field", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		field := fields.GetByName("field3")
		if field != nil {
			t.Errorf("expected nil, got %s", field.Name)
		}
	})
}
