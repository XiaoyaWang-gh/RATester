package models

import (
	"fmt"
	"testing"
)

func TestSet_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var e SmallIntegerField
	e.Set(12345)
	if e != 12345 {
		t.Errorf("Expected 12345, got %v", e)
	}
}
