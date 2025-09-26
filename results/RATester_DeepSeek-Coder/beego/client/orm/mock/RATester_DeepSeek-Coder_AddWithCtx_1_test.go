package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestAddWithCtx_1(t *testing.T) {
	type args struct {
		ctx context.Context
		i   []interface{}
	}
	tests := []struct {
		name    string
		d       *DoNothingQueryM2Mer
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

			got, err := tt.d.AddWithCtx(tt.args.ctx, tt.args.i...)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddWithCtx() got = %v, want %v", got, tt.want)
			}
		})
	}
}
