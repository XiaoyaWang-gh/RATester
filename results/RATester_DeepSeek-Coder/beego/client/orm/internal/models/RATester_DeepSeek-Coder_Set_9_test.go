package models

import (
	"fmt"
	"testing"
)

func TestSet_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var e PositiveSmallIntegerField
	d := uint16(65535)
	e.Set(d)
	if e != PositiveSmallIntegerField(d) {
		t.Errorf("Expected %v, got %v", d, e)
	}
}
