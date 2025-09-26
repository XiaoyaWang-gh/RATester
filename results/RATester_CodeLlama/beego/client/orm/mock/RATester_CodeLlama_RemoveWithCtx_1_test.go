package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestRemoveWithCtx_1(t *testing.T) {
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
		{
			name: "test1",
			d:    &DoNothingQueryM2Mer{},
			args: args{
				ctx: context.Background(),
				i:   []interface{}{},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.d.RemoveWithCtx(tt.args.ctx, tt.args.i...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingQueryM2Mer.RemoveWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingQueryM2Mer.RemoveWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
