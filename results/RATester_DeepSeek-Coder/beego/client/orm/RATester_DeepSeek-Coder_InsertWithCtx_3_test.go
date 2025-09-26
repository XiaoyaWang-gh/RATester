package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertWithCtx_3(t *testing.T) {
	type args struct {
		ctx context.Context
		md  interface{}
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

			got, err := tt.d.InsertWithCtx(tt.args.ctx, tt.args.md)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingOrm.InsertWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingOrm.InsertWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
