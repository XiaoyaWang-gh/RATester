package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	max := Max{
		Max: 10,
		Key: "test",
	}

	result := max.GetLimitValue()

	if result != 10 {
		t.Errorf("Expected 10, but got %v", result)
	}
}
