package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := Numeric{Key: "testKey"}
	expected := "testKey"
	result := n.GetKey()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
