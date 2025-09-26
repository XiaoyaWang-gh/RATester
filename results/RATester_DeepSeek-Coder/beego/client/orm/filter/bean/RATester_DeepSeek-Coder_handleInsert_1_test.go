package bean

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestHandleInsert_1(t *testing.T) {
	type args struct {
		ctx context.Context
		inv *orm.Invocation
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &DefaultValueFilterChainBuilder{}
			d.handleInsert(tt.args.ctx, tt.args.inv)
		})
	}
}
