package prometheus

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

func TestReport_1(t *testing.T) {
	type args struct {
		ctx context.Context
		inv *orm.Invocation
		dur time.Duration
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

			builder := &FilterChainBuilder{}
			builder.report(tt.args.ctx, tt.args.inv, tt.args.dur)
		})
	}
}
