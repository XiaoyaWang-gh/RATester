package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Enum{}
	result := e.GetLimitValue()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
