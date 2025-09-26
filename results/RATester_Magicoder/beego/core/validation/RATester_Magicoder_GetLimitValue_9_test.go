package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	length := Length{N: 5, Key: "testKey"}
	expected := 5
	result := length.GetLimitValue()
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
