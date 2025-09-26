package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestDeleteWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var ctx context.Context
	var result int64
	var err error

	// TODO: Add test cases.
	if result, err = o.DeleteWithCtx(ctx); err != nil {
		t.Errorf("DeleteWithCtx() error = %v, want nil", err)
		return
	}
	if result != 0 {
		t.Errorf("DeleteWithCtx() result = %v, want 0", result)
	}
}
