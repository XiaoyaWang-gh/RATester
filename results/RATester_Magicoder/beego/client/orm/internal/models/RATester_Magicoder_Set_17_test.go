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

	j := JsonbField("")
	j.Set("test")
	if j.Value() != "test" {
		t.Errorf("Expected 'test', got '%s'", j.Value())
	}
}
