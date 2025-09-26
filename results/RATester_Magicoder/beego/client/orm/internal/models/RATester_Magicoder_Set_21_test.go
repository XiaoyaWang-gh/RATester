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

	e := BooleanField(false)
	e.Set(true)
	if e != true {
		t.Errorf("Expected true, got %v", e)
	}
}
