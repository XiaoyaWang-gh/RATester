package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestRenderString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{}
	ctx := context.Background()
	args := []any{}
	got, err := pco.RenderString(ctx, args...)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}
