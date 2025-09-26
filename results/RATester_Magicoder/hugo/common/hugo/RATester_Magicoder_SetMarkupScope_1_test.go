package hugo

import (
	"context"
	"fmt"
	"testing"
)

func TestSetMarkupScope_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	scope := "test_scope"
	newCtx := SetMarkupScope(ctx, scope)

	if newCtx.Value("markupScope") != scope {
		t.Errorf("Expected markupScope to be %s, but got %v", scope, newCtx.Value("markupScope"))
	}
}
