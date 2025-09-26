package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestCountWithCtx_2(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		d       *DoNothingQuerySetter
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "case1",
			d:       &DoNothingQuerySetter{},
			args:    args{ctx: context.Background()},
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

			got, err := tt.d.CountWithCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingQuerySetter.CountWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingQuerySetter.CountWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
