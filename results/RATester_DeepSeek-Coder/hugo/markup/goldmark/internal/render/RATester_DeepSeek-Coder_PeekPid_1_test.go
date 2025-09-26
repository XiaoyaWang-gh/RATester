package render

import (
	"fmt"
	"testing"
)

func TestPeekPid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		pids: []uint64{1, 2, 3},
	}

	pid := ctx.PeekPid()
	if pid != 3 {
		t.Errorf("Expected PeekPid to return 3, got %d", pid)
	}
}
