package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	maxSize := MaxSize{Max: 10, Key: "test"}
	expected := 10
	result := maxSize.GetLimitValue()
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
