package template

import (
	"fmt"
	"testing"
)

func TestEvalArgs_1(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				args: []any{"hello"},
			},
			want: "hello",
		},
		{
			name: "case2",
			args: args{
				args: []any{"hello", "world"},
			},
			want: "helloworld",
		},
		{
			name: "case3",
			args: args{
				args: []any{"hello", "world", "!"},
			},
			want: "helloworld!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := evalArgs(tt.args.args...); got != tt.want {
				t.Errorf("evalArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
