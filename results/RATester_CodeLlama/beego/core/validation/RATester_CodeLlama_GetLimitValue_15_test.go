package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := Alpha{}
	if a.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", a.GetLimitValue())
	}
}
