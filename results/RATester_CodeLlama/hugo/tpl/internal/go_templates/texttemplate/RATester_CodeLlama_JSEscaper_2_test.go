package template

import (
	"fmt"
	"testing"
)

func TestJSEscaper_2(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				args: []any{
					"hello",
				},
			},
			want: "hello",
		},
		{
			name: "test case 2",
			args: args{
				args: []any{
					"hello",
					"world",
				},
			},
			want: "hello world",
		},
		{
			name: "test case 3",
			args: args{
				args: []any{
					"hello",
					"world",
					"!",
				},
			},
			want: "hello world !",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := JSEscaper(tt.args.args...); got != tt.want {
				t.Errorf("JSEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
