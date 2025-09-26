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

	a := AlphaNumeric{
		Key: "test",
	}

	result := a.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
