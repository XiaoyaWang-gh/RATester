package models

import (
	"fmt"
	"testing"
)

func TestSet_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := new(TextField)
	d := "test"
	e.Set(d)
	if *e != TextField(d) {
		t.Errorf("Expected %v, got %v", d, *e)
	}
}
