package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	minSize := MinSize{Min: 5, Key: "testKey"}
	expected := 5
	result := minSize.GetLimitValue()
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
