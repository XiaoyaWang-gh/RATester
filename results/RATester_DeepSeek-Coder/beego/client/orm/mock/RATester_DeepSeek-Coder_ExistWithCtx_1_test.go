package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestExistWithCtx_1(t *testing.T) {
	type args struct {
		ctx context.Context
		i   interface{}
	}
	tests := []struct {
		name string
		d    *DoNothingQueryM2Mer
		args args
		want bool
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

			if got := tt.d.ExistWithCtx(tt.args.ctx, tt.args.i); got != tt.want {
				t.Errorf("ExistWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
