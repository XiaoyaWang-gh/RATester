package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestReadOrCreateWithCtx_1(t *testing.T) {
	ctx := context.Background()
	orm := &DoNothingOrm{}

	type args struct {
		ctx  context.Context
		md   interface{}
		col1 string
		cols []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   int64
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:  ctx,
				md:   nil,
				col1: "test",
				cols: []string{"test"},
			},
			want:    false,
			want1:   0,
			wantErr: false,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1, err := orm.ReadOrCreateWithCtx(tt.args.ctx, tt.args.md, tt.args.col1, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadOrCreateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadOrCreateWithCtx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReadOrCreateWithCtx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
