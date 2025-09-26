package template

import (
	"fmt"
	"testing"
)

func TestHTMLEscaper_2(t *testing.T) {
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
					"<script>alert('hello world')</script>",
				},
			},
			want: "&lt;script&gt;alert(&#39;hello world&#39;)&lt;/script&gt;",
		},
		{
			name: "test case 2",
			args: args{
				args: []any{
					"<script>alert('hello world')</script>",
					"<script>alert('hello world')</script>",
				},
			},
			want: "&lt;script&gt;alert(&#39;hello world&#39;)&lt;/script&gt; &lt;script&gt;alert(&#39;hello world&#39;)&lt;/script&gt;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HTMLEscaper(tt.args.args...); got != tt.want {
				t.Errorf("HTMLEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
