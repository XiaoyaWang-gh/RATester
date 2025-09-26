package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestDeleteWithCtx_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o ormBase
	var ctx context.Context
	var md interface{}
	var cols []string
	var num int64
	var err error
	num, err = o.DeleteWithCtx(ctx, md, cols...)
	if err != nil {
		t.Errorf("DeleteWithCtx() error = %v", err)
		return
	}
	if num != 0 {
		t.Errorf("DeleteWithCtx() num = %v, want %v", num, 0)
	}
}
