package bean

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TesthandleInsert_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	inv := &orm.Invocation{
		Args: []interface{}{&struct{}{}},
	}
	d := &DefaultValueFilterChainBuilder{
		includeInsertOrUpdate: true,
	}
	d.handleInsert(ctx, inv)
	// assert something
}
