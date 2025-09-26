package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}
	results := &[]orm.Params{}
	exprs := []string{"expr1", "expr2"}

	_, err := d.ValuesWithCtx(ctx, results, exprs...)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
