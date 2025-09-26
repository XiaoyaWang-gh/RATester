package fiber

import (
	"fmt"
	"testing"
)

func TestgetMatched_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		matched: true,
	}

	result := ctx.getMatched()

	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}
