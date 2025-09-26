package template

import (
	"fmt"
	"testing"
)

func TestCommentEscaper_1(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testcase1",
			args: args{
				args: []any{
					"testcase1",
				},
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

			if got := commentEscaper(tt.args.args...); got != tt.want {
				t.Errorf("commentEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
