package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Email{}
	if e.GetLimitValue() != nil {
		t.Errorf("Expected nil, got %v", e.GetLimitValue())
	}
}
