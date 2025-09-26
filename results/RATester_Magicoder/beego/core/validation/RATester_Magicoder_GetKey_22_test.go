package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaNumeric{Key: "testKey"}
	expected := "testKey"
	result := a.GetKey()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
