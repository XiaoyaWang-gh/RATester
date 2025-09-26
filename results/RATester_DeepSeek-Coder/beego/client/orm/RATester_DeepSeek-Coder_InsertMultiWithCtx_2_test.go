package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertMultiWithCtx_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		bulk int
		mds  interface{}
	}
	tests := []struct {
		name    string
		d       *DoNothingOrm
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

			got, err := tt.d.InsertMultiWithCtx(tt.args.ctx, tt.args.bulk, tt.args.mds)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingOrm.InsertMultiWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingOrm.InsertMultiWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
