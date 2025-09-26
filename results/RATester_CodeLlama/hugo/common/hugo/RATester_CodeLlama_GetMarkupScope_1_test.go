package hugo

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMarkupScope_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "markupScope", "test")
	if GetMarkupScope(ctx) != "test" {
		t.Errorf("GetMarkupScope() = %v, want %v", GetMarkupScope(ctx), "test")
	}
}
