package order_clause

import (
	"fmt"
	"testing"
)

func TestSortNone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	order := &Order{}
	SortNone()(order)
	if order.sort != None {
		t.Errorf("Expected sort to be None, got %v", order.sort)
	}
}
