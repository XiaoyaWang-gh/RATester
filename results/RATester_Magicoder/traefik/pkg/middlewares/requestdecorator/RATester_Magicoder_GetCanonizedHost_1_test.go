package requestdecorator

import (
	"context"
	"fmt"
	"testing"
)

func TestGetCanonizedHost_1(t *testing.T) {
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
				ctx: context.WithValue(context.Background(), "canonicalKey", "test.com"),
			},
			want: "test.com",
		},
		{
			name: "Test case 2",
			args: args{
				ctx: context.WithValue(context.Background(), "canonicalKey", ""),
			},
			want: "",
		},
		{
			name: "Test case 3",
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

			if got := GetCanonizedHost(tt.args.ctx); got != tt.want {
				t.Errorf("GetCanonizedHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
