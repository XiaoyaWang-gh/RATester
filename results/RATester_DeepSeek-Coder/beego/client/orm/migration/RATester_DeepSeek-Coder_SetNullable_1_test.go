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

	column := &Column{
		Name:     "test_column",
		Inc:      "",
		Null:     "",
		Default:  "",
		Unsign:   "",
		DataType: "int",
		remove:   false,
		Modify:   false,
	}

	column.SetNullable(true)
	if column.Null != "" {
		t.Errorf("Expected Null to be empty string, got %s", column.Null)
	}

	column.SetNullable(false)
	if column.Null != "NOT NULL" {
		t.Errorf("Expected Null to be 'NOT NULL', got %s", column.Null)
	}
}
