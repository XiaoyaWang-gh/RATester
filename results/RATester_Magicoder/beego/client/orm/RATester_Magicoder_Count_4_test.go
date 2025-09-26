package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestCount_4(t *testing.T) {
	type args struct {
		ctx context.Context
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

			got, err := tt.o.CountWithCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryM2M.CountWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("queryM2M.CountWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
