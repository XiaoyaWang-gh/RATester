package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	min := Min{
		Min: 5,
		Key: "testKey",
	}

	expected := fmt.Sprintf(MessageTmpls["Min"], 5)
	result := min.DefaultMessage()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
