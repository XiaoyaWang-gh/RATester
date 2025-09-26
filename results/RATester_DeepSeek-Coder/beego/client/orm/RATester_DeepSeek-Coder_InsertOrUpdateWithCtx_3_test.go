package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertOrUpdateWithCtx_3(t *testing.T) {
	type args struct {
		ctx               context.Context
		md                interface{}
		colConflitAndArgs []string
	}
	tests := []struct {
		name    string
		f       *filterOrmDecorator
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

			got, err := tt.f.InsertOrUpdateWithCtx(tt.args.ctx, tt.args.md, tt.args.colConflitAndArgs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.InsertOrUpdateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("filterOrmDecorator.InsertOrUpdateWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
