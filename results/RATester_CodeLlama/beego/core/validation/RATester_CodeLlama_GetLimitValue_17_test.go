package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Required{}
	if r.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", r.GetLimitValue())
	}
}
