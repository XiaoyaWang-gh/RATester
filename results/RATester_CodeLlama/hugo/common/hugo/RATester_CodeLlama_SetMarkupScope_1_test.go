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
	ctx = SetMarkupScope(ctx, "test")
	if s := ctx.Value(markupScope).(string); s != "test" {
		t.Errorf("SetMarkupScope() = %v, want %v", s, "test")
	}
}
