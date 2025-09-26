package net

import (
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ia := &InternalAddr{}
	expected := "internal"
	actual := ia.String()
	if actual != expected {
		t.Errorf("Expected String() to return '%s', but got '%s'", expected, actual)
	}
}
