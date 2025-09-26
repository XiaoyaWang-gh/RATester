package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	max := Max{
		Max: 10,
		Key: "test",
	}
	expected := fmt.Sprintf(MessageTmpls["Max"], max.Max)
	result := max.DefaultMessage()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
