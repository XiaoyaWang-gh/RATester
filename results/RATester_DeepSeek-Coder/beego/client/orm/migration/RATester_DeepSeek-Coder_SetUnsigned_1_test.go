package migration

import (
	"fmt"
	"testing"
)

func TestSetUnsigned_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	column := &Column{
		Name: "test_column",
	}

	column.SetUnsigned(true)

	if column.Unsign != "UNSIGNED" {
		t.Errorf("Expected Unsign to be 'UNSIGNED', got %s", column.Unsign)
	}

	column.SetUnsigned(false)

	if column.Unsign != "" {
		t.Errorf("Expected Unsign to be '', got %s", column.Unsign)
	}
}
