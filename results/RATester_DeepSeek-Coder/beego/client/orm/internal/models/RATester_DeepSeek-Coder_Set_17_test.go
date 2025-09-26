package models

import (
	"fmt"
	"testing"
)

func TestSet_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := new(JsonbField)
	j.Set("test")
	if *j != "test" {
		t.Errorf("Expected 'test', got '%s'", *j)
	}
}
