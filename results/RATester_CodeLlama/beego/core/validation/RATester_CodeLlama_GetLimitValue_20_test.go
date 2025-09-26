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

	m := Min{Min: 10, Key: "key"}
	if m.GetLimitValue() != 10 {
		t.Errorf("GetLimitValue() = %v, want %v", m.GetLimitValue(), 10)
	}
}
