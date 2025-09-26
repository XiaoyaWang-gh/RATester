package hints

import (
	"fmt"
	"testing"
)

func TestDefaultRelDepth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hint := DefaultRelDepth()

	if hint.GetKey() != KeyRelDepth {
		t.Errorf("Expected KeyRelDepth, got %v", hint.GetKey())
	}

	if hint.GetValue() != true {
		t.Errorf("Expected true, got %v", hint.GetValue())
	}
}
