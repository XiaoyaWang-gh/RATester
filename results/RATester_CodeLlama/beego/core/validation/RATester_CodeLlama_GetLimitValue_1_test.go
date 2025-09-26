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

	m := MinSize{Min: 10}
	if m.GetLimitValue() != 10 {
		t.Errorf("GetLimitValue() = %v, want %v", m.GetLimitValue(), 10)
	}
}
