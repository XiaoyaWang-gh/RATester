package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Base64{}
	result := b.GetLimitValue()
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
