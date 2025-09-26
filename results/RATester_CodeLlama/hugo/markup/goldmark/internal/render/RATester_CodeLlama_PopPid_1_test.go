package render

import (
	"fmt"
	"testing"
)

func TestPopPid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		pids: []uint64{1, 2, 3},
	}
	if got := ctx.PopPid(); got != 3 {
		t.Errorf("PopPid() = %v, want %v", got, 3)
	}
	if got := ctx.PopPid(); got != 2 {
		t.Errorf("PopPid() = %v, want %v", got, 2)
	}
	if got := ctx.PopPid(); got != 1 {
		t.Errorf("PopPid() = %v, want %v", got, 1)
	}
	if got := ctx.PopPid(); got != 0 {
		t.Errorf("PopPid() = %v, want %v", got, 0)
	}
}
