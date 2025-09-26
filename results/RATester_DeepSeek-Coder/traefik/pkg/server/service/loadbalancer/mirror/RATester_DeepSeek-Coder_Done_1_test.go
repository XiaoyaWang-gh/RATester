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
	ch := ctx.Done()

	if ch == nil {
		t.Errorf("Expected a non-nil channel")
	}
}
