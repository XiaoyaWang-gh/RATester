package migration

import (
	"fmt"
	"testing"
)

func TestSetNullable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	column := &Column{}

	// Test when null is true
	column.SetNullable(true)
	if column.Null != "" {
		t.Errorf("Expected Null to be empty string, but got %s", column.Null)
	}

	// Test when null is false
	column.SetNullable(false)
	if column.Null != "NOT NULL" {
		t.Errorf("Expected Null to be 'NOT NULL', but got %s", column.Null)
	}
}
