package context

import (
	"fmt"
	"testing"
)

func TestSetParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		pnames:  []string{"key1", "key2"},
		pvalues: []string{"value1", "value2"},
	}

	input.SetParam("key1", "newValue1")

	if input.pvalues[0] != "newValue1" {
		t.Errorf("Expected pvalues[0] to be 'newValue1', but got %s", input.pvalues[0])
	}

	input.SetParam("key3", "value3")

	if len(input.pnames) != 3 || len(input.pvalues) != 3 {
		t.Errorf("Expected pnames and pvalues to have length 3, but got %d and %d", len(input.pnames), len(input.pvalues))
	}

	if input.pnames[2] != "key3" || input.pvalues[2] != "value3" {
		t.Errorf("Expected pnames[2] to be 'key3' and pvalues[2] to be 'value3', but got %s and %s", input.pnames[2], input.pvalues[2])
	}
}
