package tplimpl

import (
	"fmt"
	"testing"
)

func TestIdentifierBase_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := templateInfo{
		name: "test",
	}

	expected := "test"
	result := info.IdentifierBase()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
