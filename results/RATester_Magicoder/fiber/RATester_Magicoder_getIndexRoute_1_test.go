package fiber

import (
	"fmt"
	"testing"
)

func TestgetIndexRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		indexRoute: 1,
	}

	result := ctx.getIndexRoute()

	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
