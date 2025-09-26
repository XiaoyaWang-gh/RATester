package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFieldAlias_1(t *testing.T) {
	t.Run("Test with exported field", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		field := reflect.StructField{
			Name: "ExportedField",
			Tag:  `alias:"ef"`,
		}
		alias, options := fieldAlias(field, "alias")
		if alias != "ef" {
			t.Errorf("Expected alias to be 'ef', got '%s'", alias)
		}
		if options != nil {
			t.Errorf("Expected options to be nil, got '%v'", options)
		}
	})

	t.Run("Test with unexported field", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		field := reflect.StructField{
			Name:    "unexportedField",
			PkgPath: "main",
			Tag:     `alias:"uf"`,
		}
		alias, options := fieldAlias(field, "alias")
		if alias != "uf" {
			t.Errorf("Expected alias to be 'uf', got '%s'", alias)
		}
		if options != nil {
			t.Errorf("Expected options to be nil, got '%v'", options)
		}
	})

	t.Run("Test with no tag", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		field := reflect.StructField{
			Name: "noTagField",
		}
		alias, options := fieldAlias(field, "alias")
		if alias != "noTagField" {
			t.Errorf("Expected alias to be 'noTagField', got '%s'", alias)
		}
		if options != nil {
			t.Errorf("Expected options to be nil, got '%v'", options)
		}
	})
}
