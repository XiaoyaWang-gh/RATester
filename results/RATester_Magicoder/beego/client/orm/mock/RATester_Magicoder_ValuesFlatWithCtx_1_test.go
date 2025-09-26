package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesFlatWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}
	result := &orm.ParamsList{}
	expr := "expr"
	expectedCount := int64(0)
	count, err := d.ValuesFlatWithCtx(ctx, result, expr)
	if err != nil {
		t.Errorf("ValuesFlatWithCtx returned an error: %v", err)
	}
	if count != expectedCount {
		t.Errorf("ValuesFlatWithCtx returned count %d, expected %d", count, expectedCount)
	}
}
