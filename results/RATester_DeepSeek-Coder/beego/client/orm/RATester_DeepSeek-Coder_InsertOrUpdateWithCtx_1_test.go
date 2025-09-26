package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertOrUpdateWithCtx_1(t *testing.T) {
	type args struct {
		ctx               context.Context
		md                interface{}
		colConflitAndArgs []string
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

			got, err := tt.d.InsertOrUpdateWithCtx(tt.args.ctx, tt.args.md, tt.args.colConflitAndArgs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingOrm.InsertOrUpdateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingOrm.InsertOrUpdateWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
