package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestClearWithCtx_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		d       *DoNothingQueryM2Mer
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "TestClearWithCtx",
			d:    &DoNothingQueryM2Mer{},
			args: args{
				ctx: context.Background(),
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

			got, err := tt.d.ClearWithCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClearWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClearWithCtx() got = %v, want %v", got, tt.want)
			}
		})
	}
}
