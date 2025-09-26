package template

import (
	"fmt"
	"testing"
)

func TesthtmlEscaper_1(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				args: []any{"<script>alert('XSS')</script>"},
			},
			want: "&lt;script&gt;alert('XSS')&lt;/script&gt;",
		},
		{
			name: "Test Case 2",
			args: args{
				args: []any{"<b>bold</b>"},
			},
			want: "&lt;b&gt;bold&lt;/b&gt;",
		},
		{
			name: "Test Case 3",
			args: args{
				args: []any{"<b>bold</b>", contentTypeHTML},
			},
			want: "<b>bold</b>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := htmlEscaper(tt.args.args...); got != tt.want {
				t.Errorf("htmlEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
