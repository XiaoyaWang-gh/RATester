package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	min := Min{
		Min: 5,
		Key: "test",
	}

	expected := 5
	result := min.GetLimitValue()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
