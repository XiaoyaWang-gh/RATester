package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestBeginWithCtx_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ormer := &filterOrmDecorator{}

	txOrmer, err := ormer.BeginWithCtx(ctx)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if txOrmer == nil {
		t.Error("Expected a non-nil TxOrmer, but got nil")
	}
}
