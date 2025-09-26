package models

import (
	"fmt"
	"testing"
)

func TestSet_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := new(IntegerField)
	d := int32(12345)
	e.Set(d)
	if *e != IntegerField(d) {
		t.Errorf("Expected %v, got %v", d, *e)
	}
}
