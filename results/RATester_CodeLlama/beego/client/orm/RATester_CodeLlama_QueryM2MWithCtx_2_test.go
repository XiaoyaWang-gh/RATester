package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestQueryM2MWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var (
		d    *DoNothingOrm
		ctx  context.Context
		md   interface{}
		name string
	)
	// TODO: setup test
	// d := TODO
	// ctx := TODO
	// md := TODO
	// name := TODO
	if got := d.QueryM2MWithCtx(ctx, md, name); got != nil {
		t.Errorf("DoNothingOrm.QueryM2MWithCtx() = %v, want nil", got)
	}
}
