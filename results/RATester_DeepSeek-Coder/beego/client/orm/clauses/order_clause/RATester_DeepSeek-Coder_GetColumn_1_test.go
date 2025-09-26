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
		column: "id",
	}

	result := o.GetColumn()

	if result != "id" {
		t.Errorf("Expected 'id', got '%s'", result)
	}
}
