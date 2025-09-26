package echo

import (
	"fmt"
	"testing"
)

func TestParamValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &context{
		pvalues: []string{"value1", "value2", "value3"},
		pnames:  []string{"name1", "name2", "name3"},
	}

	result := ctx.ParamValues()

	if len(result) != len(ctx.pnames) {
		t.Errorf("Expected length of result to be %d, but got %d", len(ctx.pnames), len(result))
	}

	for i, value := range result {
		if value != ctx.pvalues[i] {
			t.Errorf("Expected value at index %d to be %s, but got %s", i, ctx.pvalues[i], value)
		}
	}
}
