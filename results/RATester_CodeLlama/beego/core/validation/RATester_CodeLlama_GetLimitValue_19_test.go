package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	z := ZipCode{}
	if z.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", z.GetLimitValue())
	}
}
