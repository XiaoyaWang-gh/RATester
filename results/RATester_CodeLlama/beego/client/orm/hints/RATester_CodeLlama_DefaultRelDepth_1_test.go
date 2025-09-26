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
		t.Errorf("Expected %v, got %v", KeyRelDepth, hint.GetKey())
	}
	if hint.GetValue() != true {
		t.Errorf("Expected %v, got %v", true, hint.GetValue())
	}
}
