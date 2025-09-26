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
	if m.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", m.GetLimitValue())
	}
}
