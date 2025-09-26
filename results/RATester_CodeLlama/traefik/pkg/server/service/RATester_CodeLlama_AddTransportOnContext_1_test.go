package service

import (
	"context"
	"fmt"
	"testing"
)

func TestAddTransportOnContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ctx = AddTransportOnContext(ctx)
	if ctx.Value(transportKey) == nil {
		t.Errorf("AddTransportOnContext failed")
	}
}
