package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestRemoveWithCtx_2(t *testing.T) {
	type args struct {
		ctx context.Context
		mds []interface{}
	}
	tests := []struct {
		name    string
		o       *queryM2M
		args    args
		want    int64
		wantErr bool
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

			got, err := tt.o.RemoveWithCtx(tt.args.ctx, tt.args.mds...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
