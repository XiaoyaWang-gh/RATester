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

	e := TextField("")
	d := "Hello, World!"
	e.Set(d)
	if e != TextField(d) {
		t.Errorf("Expected %s, got %s", d, e)
	}
}
