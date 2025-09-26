package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaNumeric{Key: "testKey"}
	expected := "AlphaNumeric"
	actual := a.DefaultMessage()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
