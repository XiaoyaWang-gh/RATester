package template

import (
	"fmt"
	"testing"
)

func TestHtmlEscaper_1(t *testing.T) {
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
			args: args{args: []any{"<script>alert('hello')</script>"}},
			want: "&lt;script&gt;alert(&#39;hello&#39;)&lt;/script&gt;",
		},
		{
			name: "test2",
			args: args{args: []any{"<script>alert('hello')</script>", contentTypeHTML}},
			want: "<script>alert('hello')</script>",
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
