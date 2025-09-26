package context

import (
	"fmt"
	"testing"
)

func TestReset_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	ctx := &Context{}
	input.Reset(ctx)
	if input.Context != ctx {
		t.Errorf("input.Context should be ctx")
	}
	if input.CruSession != nil {
		t.Errorf("input.CruSession should be nil")
	}
	if len(input.pnames) != 0 {
		t.Errorf("input.pnames should be empty")
	}
	if len(input.pvalues) != 0 {
		t.Errorf("input.pvalues should be empty")
	}
	if input.data != nil {
		t.Errorf("input.data should be nil")
	}
	if len(input.RequestBody) != 0 {
		t.Errorf("input.RequestBody should be empty")
	}
}
