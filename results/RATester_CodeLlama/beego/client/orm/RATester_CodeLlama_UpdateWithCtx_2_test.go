package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestUpdateWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o ormBase
	var ctx context.Context
	var md interface{}
	var cols []string
	var want int64
	var err error
	var got int64
	var gotErr error

	// TODO

	if got, gotErr = o.UpdateWithCtx(ctx, md, cols...); (err != nil) != (gotErr != nil) {
		t.Errorf("UpdateWithCtx() error = %v, wantErr %v", gotErr, err)
		return
	}
	if got != want {
		t.Errorf("UpdateWithCtx() got = %v, want %v", got, want)
	}
}
