package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestUpdateWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}
	values := orm.Params{"key": "value"}

	_, err := d.UpdateWithCtx(ctx, values)
	if err != nil {
		t.Errorf("UpdateWithCtx returned an error: %v", err)
	}
}
