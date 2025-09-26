package fiber

import (
	"context"
	"fmt"
	"testing"
)

func TestSetUserContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{}
	userCtx := context.Background()
	ctx.SetUserContext(userCtx)

	result := ctx.UserContext()
	if result != userCtx {
		t.Errorf("Expected %v, got %v", userCtx, result)
	}
}
