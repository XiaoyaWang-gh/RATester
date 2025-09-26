package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestredirectFixedPath_1(t *testing.T) {
	type args struct {
		c             *Context
		root          *node
		trailingSlash bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case 1",
			args: args{
				c: &Context{
					Request: &http.Request{
						URL: &url.URL{
							Path: "/test",
						},
					},
				},
				root: &node{
					children: []*node{
						{
							path: "/test",
						},
					},
				},
				trailingSlash: true,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				c: &Context{
					Request: &http.Request{
						URL: &url.URL{
							Path: "/test",
						},
					},
				},
				root: &node{
					children: []*node{
						{
							path: "/test2",
						},
					},
				},
				trailingSlash: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := redirectFixedPath(tt.args.c, tt.args.root, tt.args.trailingSlash); got != tt.want {
				t.Errorf("redirectFixedPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
