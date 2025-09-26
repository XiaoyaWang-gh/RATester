package orm

import (
	"fmt"
	"testing"
)

func TestOrderBy_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{}

	// Test with valid expressions
	qs.OrderBy("column1", "-column2")

	// Test with no expressions
	qs.OrderBy()

	// Test with invalid expressions
	qs.OrderBy("invalid_column")
}
