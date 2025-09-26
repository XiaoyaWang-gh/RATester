package order_clause

import (
	"fmt"
	"testing"
)

func TestIsRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Order{}
	if got := o.IsRaw(); got != false {
		t.Errorf("IsRaw() = %v, want %v", got, false)
	}
}
