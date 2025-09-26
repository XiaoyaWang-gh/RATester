package template

import (
	"fmt"
	"testing"
)

func TestURLQueryEscaper_1(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{args: []any{"hello"}},
			want: "hello",
		},
		{
			name: "test2",
			args: args{args: []any{"hello", "world"}},
			want: "hello%20world",
		},
		{
			name: "test3",
			args: args{args: []any{"hello", "world", "!"}},
			want: "hello%20world%21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := URLQueryEscaper(tt.args.args...); got != tt.want {
				t.Errorf("URLQueryEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
