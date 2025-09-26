package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertMultiWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var (
		f    *filterOrmDecorator
		ctx  context.Context
		bulk int
		mds  interface{}
	)

	f.InsertMultiWithCtx(ctx, bulk, mds)
}
