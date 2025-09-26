package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestBeginWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	orm := &DoNothingOrm{}

	tx, err := orm.BeginWithCtx(ctx)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if tx != nil {
		t.Errorf("Expected nil transaction, but got %v", tx)
	}
}
