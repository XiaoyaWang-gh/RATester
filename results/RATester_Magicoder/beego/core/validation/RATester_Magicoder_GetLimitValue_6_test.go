package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaNumeric{Key: "testKey"}
	result := a.GetLimitValue()
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
