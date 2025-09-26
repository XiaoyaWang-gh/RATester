package models

import (
	"fmt"
	"testing"
)

func TestSet_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var e PositiveBigIntegerField
	d := uint64(18446744073709551615)
	e.Set(d)
	if e != PositiveBigIntegerField(d) {
		t.Errorf("Expected %v, got %v", d, e)
	}
}
