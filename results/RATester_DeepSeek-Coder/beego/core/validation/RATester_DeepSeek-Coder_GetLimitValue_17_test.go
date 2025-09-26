package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Required{
		Key: "testKey",
	}

	result := r.GetLimitValue()

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
