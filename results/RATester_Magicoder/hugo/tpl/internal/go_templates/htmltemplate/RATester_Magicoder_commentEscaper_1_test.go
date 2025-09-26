package template

import (
	"fmt"
	"testing"
)

func TestcommentEscaper_1(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				args: []any{"field name"},
			},
			want: "field want",
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
