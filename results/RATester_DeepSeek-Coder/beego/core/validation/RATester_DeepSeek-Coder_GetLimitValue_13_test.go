package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Mobile{}
	result := m.GetLimitValue()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
