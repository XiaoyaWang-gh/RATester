package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestFilterChain_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &OrmStub{}
	next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
		return []interface{}{"next"}
	}
	f := o.FilterChain(next)
	ctx := context.Background()
	inv := &orm.Invocation{}
	res := f(ctx, inv)
	if len(res) != 1 {
		t.Errorf("FilterChain() = %v, want %v", res, 1)
	}
	if res[0] != "next" {
		t.Errorf("FilterChain() = %v, want %v", res[0], "next")
	}
}
