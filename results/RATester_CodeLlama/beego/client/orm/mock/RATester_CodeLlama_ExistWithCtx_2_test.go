package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestExistWithCtx_2(t *testing.T) {
	type fields struct {
		DoNothingQuerySetter *DoNothingQuerySetter
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "TestExistWithCtx",
			fields: fields{DoNothingQuerySetter: &DoNothingQuerySetter{}},
			args:   args{ctx: context.Background()},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &DoNothingQuerySetter{}
			if got := d.ExistWithCtx(tt.args.ctx); got != tt.want {
				t.Errorf("ExistWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
