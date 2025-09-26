package order_clause

import (
	"fmt"
	"testing"
)

func TestGetColumn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Order{
		column: "test_column",
	}

	result := o.GetColumn()

	if result != "test_column" {
		t.Errorf("Expected 'test_column', but got '%s'", result)
	}
}
