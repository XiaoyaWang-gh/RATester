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

	c := &Context{
		index: 0,
	}
	c.Abort()
	if c.index != abortIndex {
		t.Errorf("Abort() failed")
	}
}
