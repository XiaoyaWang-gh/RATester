package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Phone{}
	if p.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", p.GetLimitValue())
	}
}
