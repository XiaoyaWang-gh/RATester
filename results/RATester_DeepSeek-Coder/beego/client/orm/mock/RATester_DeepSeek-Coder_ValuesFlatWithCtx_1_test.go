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

	t.Parallel()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}
	result := &orm.ParamsList{}
	expr := "test"

	_, err := d.ValuesFlatWithCtx(ctx, result, expr)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	cancel()
	_, err = d.ValuesFlatWithCtx(ctx, result, expr)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
