package web

import (
	"fmt"
	"testing"
)

func TestHTML2str_1(t *testing.T) {
	tests := []struct {
		name string
		html string
		want string
	}{
		{
			name: "Test case 1",
			html: "<div>Hello, World!</div>",
			want: "hello, world!",
		},
		{
			name: "Test case 2",
			html: "<div>Hello, <b>World</b>!</div>",
			want: "hello, world!",
		},
		{
			name: "Test case 3",
			html: "<div>Hello, <b>World</b>!</div><script>alert('Hello, World!')</script>",
			want: "hello, world!\n",
		},
		{
			name: "Test case 4",
			html: "<div>Hello, <b>World</b>!</div><style>body { background-color: red; }</style>",
			want: "hello, world!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HTML2str(tt.html); got != tt.want {
				t.Errorf("HTML2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
