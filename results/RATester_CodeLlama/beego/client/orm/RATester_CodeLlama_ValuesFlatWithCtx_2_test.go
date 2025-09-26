package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestValuesFlatWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var ctx context.Context
	var result *ParamsList
	var expr string
	var err error
	var i int64

	i, err = o.ValuesFlatWithCtx(ctx, result, expr)
	if err != nil {
		t.Fatal(err)
	}
	if i != 0 {
		t.Fatalf("Expected 0, got %v", i)
	}
}
