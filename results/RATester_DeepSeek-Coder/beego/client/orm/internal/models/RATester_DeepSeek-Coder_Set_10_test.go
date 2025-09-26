package models

import (
	"fmt"
	"testing"
)

func TestSet_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := new(CharField)
	d := "test"
	e.Set(d)
	if *e != CharField(d) {
		t.Errorf("Expected %v, got %v", d, *e)
	}
}
