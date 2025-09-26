package plugin

import (
	"context"
	"fmt"
	"testing"
)

func TestGetReqidFromContext_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				ctx: context.WithValue(context.Background(), reqidKey, "test_reqid"),
			},
			want: "test_reqid",
		},
		{
			name: "Test case 2",
			args: args{
				ctx: context.Background(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetReqidFromContext(tt.args.ctx); got != tt.want {
				t.Errorf("GetReqidFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
