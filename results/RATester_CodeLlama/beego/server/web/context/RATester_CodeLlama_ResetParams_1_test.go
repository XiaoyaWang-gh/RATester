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

	input := &BeegoInput{}
	input.pnames = []string{"a", "b"}
	input.pvalues = []string{"c", "d"}
	input.ResetParams()
	if len(input.pnames) != 0 {
		t.Errorf("pnames should be empty")
	}
	if len(input.pvalues) != 0 {
		t.Errorf("pvalues should be empty")
	}
}
