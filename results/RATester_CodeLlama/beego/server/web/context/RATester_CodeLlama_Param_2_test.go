package context

import (
	"fmt"
	"testing"
)

func TestParam_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		pnames:  []string{"key1", "key2"},
		pvalues: []string{"value1", "value2"},
	}
	key := "key1"
	if input.Param(key) != "value1" {
		t.Errorf("Param() = %v, want %v", input.Param(key), "value1")
	}
}
