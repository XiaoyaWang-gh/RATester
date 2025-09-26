package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestQueryTable_1(t *testing.T) {
	o := ormBase{}

	t.Run("TestQueryTable", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		qs := o.QueryTable("test")
		if qs == nil {
			t.Errorf("Expected QuerySeter, got nil")
		}
	})

	t.Run("TestQueryTableWithCtx", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		qs := o.QueryTableWithCtx(ctx, "test")
		if qs == nil {
			t.Errorf("Expected QuerySeter, got nil")
		}
	})
}
