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

	e := SmallIntegerField(0)
	d := int16(100)
	e.Set(d)
	if e != SmallIntegerField(d) {
		t.Errorf("Expected %v, got %v", d, e)
	}
}
