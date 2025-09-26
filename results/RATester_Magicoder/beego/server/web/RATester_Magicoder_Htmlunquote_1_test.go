package web

import (
	"fmt"
	"testing"
)

func TestHtmlunquote_1(t *testing.T) {

	type args struct {
		text string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{text: "&lt;p&gt;Hello, World!&lt;/p&gt;"},
			want: "<p>Hello, World!</p>",
		},
		{
			name: "Test Case 2",
			args: args{text: "&amp;lt;p&amp;gt;Hello, World!&amp;lt;/p&amp;gt;"},
			want: "&lt;p&gt;Hello, World!&lt;/p&gt;",
		},
		{
			name: "Test Case 3",
			args: args{text: "&lt;p&gt;Hello, World!&lt;/p&gt; "},
			want: "<p>Hello, World!</p>",
		},
		{
			name: "Test Case 4",
			args: args{text: " &lt;p&gt;Hello, World!&lt;/p&gt; "},
			want: "<p>Hello, World!</p>",
		},
		{
			name: "Test Case 5",
			args: args{text: "&lt;p&gt;Hello, World!&lt;/p&gt; "},
			want: "<p>Hello, World!</p>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Htmlunquote(tt.args.text); got != tt.want {
				t.Errorf("Htmlunquote() = %v, want %v", got, tt.want)
			}
		})
	}
}
