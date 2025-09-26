package order_clause

import (
	"fmt"
	"testing"
)

func TestColumn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var order Order
	column := "column"
	option := Column(column)
	option(&order)
	if order.GetColumn() != column {
		t.Errorf("Column() failed")
	}
}
