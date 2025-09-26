package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := Alpha{Key: "testKey"}
	expected := "Alpha"
	actual := a.DefaultMessage()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
