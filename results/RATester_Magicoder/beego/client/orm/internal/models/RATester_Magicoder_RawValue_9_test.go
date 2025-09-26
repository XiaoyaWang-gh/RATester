package models

import (
	"fmt"
	"testing"
)

func TestRawValue_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := JSONField("test")
	expected := "test"
	result := j.RawValue()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
