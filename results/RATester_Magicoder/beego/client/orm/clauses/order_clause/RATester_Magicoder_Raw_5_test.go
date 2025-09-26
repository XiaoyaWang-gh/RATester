package order_clause

import (
	"fmt"
	"testing"
)

func TestRaw_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	order := &Order{}
	Raw()(order)

	if !order.isRaw {
		t.Error("Expected isRaw to be true, but it was false")
	}
}
