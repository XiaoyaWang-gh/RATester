package context

import (
	"fmt"
	"testing"
)

func TestResetParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		pnames:  []string{"name", "age"},
		pvalues: []string{"John", "30"},
	}

	input.ResetParams()

	if len(input.pnames) != 0 || len(input.pvalues) != 0 {
		t.Errorf("Expected pnames and pvalues to be empty after ResetParams, but got %v and %v", input.pnames, input.pvalues)
	}
}
