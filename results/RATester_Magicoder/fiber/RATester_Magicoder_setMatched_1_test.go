package fiber

import (
	"fmt"
	"testing"
)

func TestsetMatched_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{}
	ctx.setMatched(true)
	if ctx.matched != true {
		t.Errorf("Expected matched to be true, but got %v", ctx.matched)
	}
}
