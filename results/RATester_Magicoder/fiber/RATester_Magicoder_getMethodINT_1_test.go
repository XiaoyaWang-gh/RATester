package fiber

import (
	"fmt"
	"testing"
)

func TestgetMethodINT_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		methodINT: 1,
	}

	result := ctx.getMethodINT()

	if result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}
}
