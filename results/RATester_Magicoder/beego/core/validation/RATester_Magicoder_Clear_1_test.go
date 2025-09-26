package validation

import (
	"fmt"
	"testing"
)

func TestClear_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	v.Clear()
	if len(v.Errors) != 0 {
		t.Errorf("Expected Errors to be empty, but got %v", v.Errors)
	}
	if v.ErrorsMap != nil {
		t.Errorf("Expected ErrorsMap to be nil, but got %v", v.ErrorsMap)
	}
}
