package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaDash{}
	if got := a.GetLimitValue(); got != nil {
		t.Errorf("AlphaDash.GetLimitValue() = %v, want nil", got)
	}
}
