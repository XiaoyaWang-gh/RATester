package types

import (
	"fmt"
	"testing"
)

func TestPrintableValue_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Result[int]{Value: 1}
	if got := r.PrintableValue(); got != 1 {
		t.Errorf("PrintableValue() = %v, want %v", got, 1)
	}
}
