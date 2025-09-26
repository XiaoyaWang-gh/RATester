package opentracing

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestFilterChain_11(t *testing.T) {
	type args struct {
		ctx  context.Context
		inv  *orm.Invocation
		next orm.Filter
	}
	tests := []struct {
		name string
		args args
		want []interface{}
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

			builder := &FilterChainBuilder{}
			if got := builder.FilterChain(tt.args.next)(tt.args.ctx, tt.args.inv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
