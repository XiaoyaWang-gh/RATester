package xlog

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestFromContextSafe_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *Logger
	}{
		{
			name: "Test case 1",
			args: args{
				ctx: context.Background(),
			},
			want: &Logger{},
		},
		{
			name: "Test case 2",
			args: args{
				ctx: context.WithValue(context.Background(), xlogKey, &Logger{}),
			},
			want: &Logger{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FromContextSafe(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromContextSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
