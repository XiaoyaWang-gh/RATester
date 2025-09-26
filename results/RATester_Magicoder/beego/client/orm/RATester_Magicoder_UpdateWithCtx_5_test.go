package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
)

func TestUpdateWithCtx_5(t *testing.T) {
	ctx := context.Background()
	orm := &DoNothingOrm{}

	type args struct {
		ctx  context.Context
		md   interface{}
		cols []string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:  ctx,
				md:   &model.User{},
				cols: []string{"name"},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				ctx:  ctx,
				md:   &model.User{},
				cols: []string{"email"},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				ctx:  ctx,
				md:   &model.User{},
				cols: []string{"password"},
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

			got, err := orm.UpdateWithCtx(tt.args.ctx, tt.args.md, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
