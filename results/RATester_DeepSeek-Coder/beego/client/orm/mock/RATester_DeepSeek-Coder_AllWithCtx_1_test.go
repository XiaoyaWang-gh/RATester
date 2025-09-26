package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestAllWithCtx_1(t *testing.T) {
	type args struct {
		ctx       context.Context
		container interface{}
		cols      []string
	}
	tests := []struct {
		name    string
		d       *DoNothingQuerySetter
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

			got, err := tt.d.AllWithCtx(tt.args.ctx, tt.args.container, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AllWithCtx() got = %v, want %v", got, tt.want)
			}
		})
	}
}
