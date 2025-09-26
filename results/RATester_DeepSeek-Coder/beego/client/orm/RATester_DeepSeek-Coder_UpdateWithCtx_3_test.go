package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestUpdateWithCtx_3(t *testing.T) {
	type args struct {
		ctx    context.Context
		values Params
	}
	tests := []struct {
		name    string
		o       querySet
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

			got, err := tt.o.UpdateWithCtx(tt.args.ctx, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("querySet.UpdateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("querySet.UpdateWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
