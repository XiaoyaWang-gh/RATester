package mirror

import (
	"fmt"
	"testing"
)

func TestDone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := contextStopPropagation{}
	done := ctx.Done()
	if done == nil {
		t.Error("Expected a channel, got nil")
	}
}
