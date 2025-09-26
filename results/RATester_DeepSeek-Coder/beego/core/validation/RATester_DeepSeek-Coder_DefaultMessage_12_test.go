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

	minSize := MinSize{Min: 5}
	expected := fmt.Sprintf(MessageTmpls["MinSize"], 5)
	result := minSize.DefaultMessage()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
