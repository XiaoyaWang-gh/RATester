package models

import (
	"fmt"
	"testing"
)

func TestSet_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := new(BooleanField)
	e.Set(true)
	if *e != true {
		t.Errorf("Expected true, got %v", *e)
	}
	e.Set(false)
	if *e != false {
		t.Errorf("Expected false, got %v", *e)
	}
}
