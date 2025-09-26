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

	c := &DefaultCtx{}
	ctx := context.Background()
	c.SetUserContext(ctx)
	if c.UserContext() != ctx {
		t.Errorf("SetUserContext() = %v, want %v", c.UserContext(), ctx)
	}
}
