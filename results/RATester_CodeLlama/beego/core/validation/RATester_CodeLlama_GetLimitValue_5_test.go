package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := IP{}
	if i.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", i.GetLimitValue())
	}
}
