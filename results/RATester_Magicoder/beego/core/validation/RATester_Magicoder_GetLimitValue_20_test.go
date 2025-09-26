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

	min := Min{Min: 10, Key: "testKey"}
	expected := 10
	result := min.GetLimitValue()
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
