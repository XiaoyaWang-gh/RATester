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

	e := CharField("")
	d := "test"
	e.Set(d)
	if string(e) != d {
		t.Errorf("Expected %s, got %s", d, e)
	}
}
