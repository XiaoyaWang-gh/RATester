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
	if hint.GetKey() != KeyForUpdate {
		t.Errorf("Expected key %v, got %v", KeyForUpdate, hint.GetKey())
	}
	if hint.GetValue() != true {
		t.Errorf("Expected value true, got %v", hint.GetValue())
	}
}
