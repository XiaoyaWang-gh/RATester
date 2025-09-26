package provider

import (
	"context"
	"fmt"
	"testing"
)

func TestAddInContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ctx = AddInContext(ctx, "test")
	if ctx.Value(key).(string) != "test" {
		t.Errorf("AddInContext failed")
	}
}
