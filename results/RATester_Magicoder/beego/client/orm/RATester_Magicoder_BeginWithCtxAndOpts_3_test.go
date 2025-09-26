package orm

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestBeginWithCtxAndOpts_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	opts := &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}
	orm := &DoNothingOrm{}
	txOrm, err := orm.BeginWithCtxAndOpts(ctx, opts)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if txOrm == nil {
		t.Error("Expected non-nil TxOrmer, got nil")
	}
}
