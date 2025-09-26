package bean

import (
	"context"
	"fmt"
	"testing"
)

func TestSetDefaultValue_1(t *testing.T) {
	type args struct {
		ctx context.Context
		ins interface{}
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
			d.setDefaultValue(tt.args.ctx, tt.args.ins)
		})
	}
}
