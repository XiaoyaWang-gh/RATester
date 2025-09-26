package internal

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestnodeContent_1(t *testing.T) {
	tests := []struct {
		name string
		node *html.Node
		want string
	}{
		{
			name: "Test Case 1",
			node: &html.Node{
				Type: html.ElementNode,
				Data: "div",
				FirstChild: &html.Node{
					Type: html.TextNode,
					Data: "Hello, World!",
				},
			},
			want: "Hello, World!",
		},
		{
			name: "Test Case 2",
			node: &html.Node{
				Type: html.ElementNode,
				Data: "p",
				FirstChild: &html.Node{
					Type: html.TextNode,
					Data: "This is a test.",
				},
			},
			want: "This is a test.",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := nodeContent(tt.node); got != tt.want {
				t.Errorf("nodeContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
