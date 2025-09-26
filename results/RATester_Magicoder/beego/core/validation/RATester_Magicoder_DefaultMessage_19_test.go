package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Required{Key: "testKey"}
	expected := "testKey is required"
	actual := r.DefaultMessage()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
