package models

import (
	"fmt"
	"testing"
)

func TestValue_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := JsonbField("test")
	expected := "test"
	result := j.Value()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
