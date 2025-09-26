package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestReadWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	orm := &DoNothingOrm{}

	// Test case 1: Successful read
	md := &struct{}{}
	err := orm.ReadWithCtx(ctx, md)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Error in read
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	err = orm.ReadWithCtx(ctx, md)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
