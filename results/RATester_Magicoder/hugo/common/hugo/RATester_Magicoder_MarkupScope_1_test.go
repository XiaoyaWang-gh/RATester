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

	ctx := context.Background()
	c := Context{}

	expected := "expected_value"
	ctx = context.WithValue(ctx, "markup_scope", expected)

	result := c.MarkupScope(ctx)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
