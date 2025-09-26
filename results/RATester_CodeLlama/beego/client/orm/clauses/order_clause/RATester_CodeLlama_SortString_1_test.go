package order_clause

import (
	"fmt"
	"testing"
)

func TestSortString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Order{
		column: "id",
		sort:   Ascending,
	}

	if o.SortString() != "ASC" {
		t.Errorf("Expected %s, got %s", "ASC", o.SortString())
	}
}
