package tplimpl

import (
	"fmt"
	"testing"
)

func TestIsZero_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := templateInfo{
		name: "",
	}

	if !info.IsZero() {
		t.Errorf("Expected IsZero to return true, but it returned false")
	}

	info.name = "test"
	if info.IsZero() {
		t.Errorf("Expected IsZero to return false, but it returned true")
	}
}
