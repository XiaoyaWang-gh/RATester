package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := Numeric{Key: "testKey"}
	expected := "Numeric check number"
	actual := n.DefaultMessage()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
