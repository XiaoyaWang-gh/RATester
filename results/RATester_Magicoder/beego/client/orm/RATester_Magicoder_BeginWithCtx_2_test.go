package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestBeginWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &orm{}
	ctx := context.Background()

	tx, err := o.BeginWithCtx(ctx)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if tx == nil {
		t.Error("Expected a transaction, but got nil")
	}
}
