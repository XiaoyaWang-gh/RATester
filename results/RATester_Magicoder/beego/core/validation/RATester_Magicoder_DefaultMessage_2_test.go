package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	maxSize := MaxSize{Max: 10, Key: "testKey"}
	expected := fmt.Sprintf(MessageTmpls["MaxSize"], maxSize.Max)
	result := maxSize.DefaultMessage()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
