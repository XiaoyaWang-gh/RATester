package validation

import (
	"fmt"
	"testing"
)

func TestGetLimitValue_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := Length{N: 10}
	if l.GetLimitValue() != 10 {
		t.Errorf("GetLimitValue() = %v, want %v", l.GetLimitValue(), 10)
	}
}
