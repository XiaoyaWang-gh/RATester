package template

import (
	"fmt"
	"testing"
)

func TestJsStrEscaper_1(t *testing.T) {
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
			args: args{args: []any{"test1"}},
			want: "test1",
		},
		{
			name: "test2",
			args: args{args: []any{"test2"}},
			want: "test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := jsStrEscaper(tt.args.args...); got != tt.want {
				t.Errorf("jsStrEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
