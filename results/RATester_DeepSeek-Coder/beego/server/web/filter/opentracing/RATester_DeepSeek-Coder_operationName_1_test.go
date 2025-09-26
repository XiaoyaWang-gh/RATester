package opentracing

import (
	"fmt"
	"testing"

	beegoCtx "github.com/beego/beego/v2/server/web/context"
)

func TestOperationName_1(t *testing.T) {
	type args struct {
		ctx *beegoCtx.Context
	}
	tests := []struct {
		name string
		args args
		want string
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
			if got := builder.operationName(tt.args.ctx); got != tt.want {
				t.Errorf("operationName() = %v, want %v", got, tt.want)
			}
		})
	}
}
