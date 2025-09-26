package failover

import (
	"context"
	"fmt"
	"testing"
)

func TestSetFallbackHandlerStatus_1(t *testing.T) {
	type args struct {
		ctx context.Context
		up  bool
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

			f := &Failover{}
			f.SetFallbackHandlerStatus(tt.args.ctx, tt.args.up)
		})
	}
}
