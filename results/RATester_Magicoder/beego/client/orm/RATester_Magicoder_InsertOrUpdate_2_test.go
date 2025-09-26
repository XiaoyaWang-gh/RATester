package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertOrUpdate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &struct{}{}
	colConflictAndArgs := []string{"colConflict", "arg1", "arg2"}
	ctx := context.Background()

	_, err := o.InsertOrUpdate(md, colConflictAndArgs...)
	if err != nil {
		t.Errorf("InsertOrUpdate() error = %v", err)
		return
	}

	_, err = o.InsertOrUpdateWithCtx(ctx, md, colConflictAndArgs...)
	if err != nil {
		t.Errorf("InsertOrUpdateWithCtx() error = %v", err)
		return
	}
}
