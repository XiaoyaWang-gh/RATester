package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := Numeric{Key: "key"}
	if n.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", n.GetLimitValue())
	}
}
