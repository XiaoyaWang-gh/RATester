package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Enum{
		Rules: "enum=1,2,3",
		Key:   "enum",
	}
	if e.GetLimitValue() != nil {
		t.Errorf("GetLimitValue() = %v, want nil", e.GetLimitValue())
	}
}
