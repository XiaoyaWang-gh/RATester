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
	if hint == nil {
		t.Error("Expected a non-nil hint, but got nil")
	}
	if hint.GetKey() != KeyRelDepth {
		t.Errorf("Expected key to be %v, but got %v", KeyRelDepth, hint.GetKey())
	}
	if hint.GetValue() != true {
		t.Errorf("Expected value to be true, but got %v", hint.GetValue())
	}
}
