package hints

import (
	"fmt"
	"testing"
)

func TestForUpdate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hint := ForUpdate()
	if hint == nil {
		t.Error("Expected a non-nil hint, but got nil")
	}

	if hint.GetKey() != KeyForUpdate {
		t.Errorf("Expected key %v, but got %v", KeyForUpdate, hint.GetKey())
	}

	if hint.GetValue() != true {
		t.Errorf("Expected value true, but got %v", hint.GetValue())
	}
}
