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

	t.Parallel()

	ctx := context.Background()
	c := Context{}

	result := c.MarkupScope(ctx)

	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}
