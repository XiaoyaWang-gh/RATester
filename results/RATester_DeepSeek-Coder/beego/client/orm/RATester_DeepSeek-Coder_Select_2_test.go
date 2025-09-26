package orm

import (
	"fmt"
	"testing"
)

func TestSelect_2(t *testing.T) {
	qb := &PostgresQueryBuilder{}

	t.Run("Select all fields", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qb.Select("*")
		expected := "SELECT *"
		result := qb.String()
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Select specific fields", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qb.Select("field1", "field2")
		expected := "SELECT \"field1\",\"field2\""
		result := qb.String()
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Select fields with table prefix", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qb.Select("table1.field1", "table2.field2")
		expected := "SELECT \"table1\".\"field1\",\"table2\".\"field2\""
		result := qb.String()
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
