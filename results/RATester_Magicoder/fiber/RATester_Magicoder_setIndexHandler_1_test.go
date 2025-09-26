package fiber

import (
	"fmt"
	"testing"
)

func TestsetIndexHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{}
	handler := 10
	ctx.setIndexHandler(handler)
	if ctx.indexHandler != handler {
		t.Errorf("Expected %d, got %d", handler, ctx.indexHandler)
	}
}
