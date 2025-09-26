package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaNumeric{Key: "key"}
	if a.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", a.GetLimitValue())
	}
}
