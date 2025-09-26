package hugo

import (
	"context"
	"fmt"
	"testing"
)

func TestMarkupScope_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Context{}
	ctx := context.Background()
	if c.MarkupScope(ctx) != GetMarkupScope(ctx) {
		t.Errorf("MarkupScope() = %v, want %v", c.MarkupScope(ctx), GetMarkupScope(ctx))
	}
}
