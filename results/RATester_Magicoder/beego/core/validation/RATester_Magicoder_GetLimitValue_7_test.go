package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := Numeric{Key: "testKey"}
	result := n.GetLimitValue()
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
