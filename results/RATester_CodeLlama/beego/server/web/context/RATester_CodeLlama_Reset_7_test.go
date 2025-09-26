package context

import (
	"fmt"
	"testing"
)

func TestReset_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	ctx := &Context{}
	output.Reset(ctx)
	if output.Context != ctx {
		t.Errorf("output.Context should be ctx")
	}
	if output.Status != 0 {
		t.Errorf("output.Status should be 0")
	}
}
