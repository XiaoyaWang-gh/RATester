package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Max{Max: 10, Key: "key"}
	if m.GetLimitValue() != 10 {
		t.Errorf("GetLimitValue() = %v, want %v", m.GetLimitValue(), 10)
	}
}
