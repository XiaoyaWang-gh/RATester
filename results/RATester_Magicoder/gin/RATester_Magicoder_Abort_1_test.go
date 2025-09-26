package gin

import (
	"fmt"
	"testing"
)

func TestAbort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	context := &Context{
		index: 0,
	}

	context.Abort()

	if context.index != abortIndex {
		t.Errorf("Expected index to be %d, but got %d", abortIndex, context.index)
	}
}
