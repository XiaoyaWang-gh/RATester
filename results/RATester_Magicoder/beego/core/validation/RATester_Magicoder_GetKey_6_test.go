package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := NoMatch{
		Key: "testKey",
	}

	expected := "testKey"
	result := n.GetKey()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
