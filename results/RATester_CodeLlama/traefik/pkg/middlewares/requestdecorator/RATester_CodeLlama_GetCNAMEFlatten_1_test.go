package requestdecorator

import (
	"context"
	"fmt"
	"testing"
)

func TestGetCNAMEFlatten_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ctx = context.WithValue(ctx, flattenKey, "flatten")
	if GetCNAMEFlatten(ctx) != "flatten" {
		t.Errorf("GetCNAMEFlatten() = %v, want %v", GetCNAMEFlatten(ctx), "flatten")
	}
}
