package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestExistWithCtx_3(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		o    querySet
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

			if got := tt.o.ExistWithCtx(tt.args.ctx); got != tt.want {
				t.Errorf("querySet.ExistWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
