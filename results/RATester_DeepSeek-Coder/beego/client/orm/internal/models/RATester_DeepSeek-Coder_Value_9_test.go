package models

import (
	"fmt"
	"testing"
)

func TestValue_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := IntegerField(12345)
	expected := int32(12345)
	result := e.Value()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
