package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Enum{
		Rules: "test",
		Key:   "test",
	}

	expected := fmt.Sprintf(MessageTmpls["Enum"], e.Rules)
	actual := e.DefaultMessage()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
