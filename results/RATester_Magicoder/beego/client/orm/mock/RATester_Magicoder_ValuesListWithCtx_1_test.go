package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesListWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}
	results := []orm.ParamsList{}
	_, err := d.ValuesListWithCtx(ctx, &results, "col1", "col2")
	if err != nil {
		t.Errorf("ValuesListWithCtx() error = %v, wantErr %v", err, false)
		return
	}
}
