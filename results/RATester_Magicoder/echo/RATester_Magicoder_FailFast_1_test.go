package echo

import (
	"fmt"
	"testing"
)

func TestFailFast_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{}
	expected := true
	result := b.FailFast(expected)
	if result.failFast != expected {
		t.Errorf("Expected failFast to be %v, but got %v", expected, result.failFast)
	}
}
