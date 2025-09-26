package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	minSize := MinSize{Min: 5, Key: "testKey"}
	expected := fmt.Sprintf(MessageTmpls["MinSize"], minSize.Min)
	result := minSize.DefaultMessage()
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
