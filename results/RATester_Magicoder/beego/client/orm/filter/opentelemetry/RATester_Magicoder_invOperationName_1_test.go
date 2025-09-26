package opentelemetry

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestinvOperationName_1(t *testing.T) {
	type args struct {
		ctx context.Context
		inv *orm.Invocation
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

			if got := invOperationName(tt.args.ctx, tt.args.inv); got != tt.want {
				t.Errorf("invOperationName() = %v, want %v", got, tt.want)
			}
		})
	}
}
