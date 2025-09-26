package internal

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestAttr_1(t *testing.T) {
	tests := []struct {
		name string
		node *html.Node
		key  string
		want string
	}{
		{
			name: "Test case 1",
			node: &html.Node{
				Type: html.ElementNode,
				Data: "div",
				Attr: []html.Attribute{
					{Key: "id", Val: "test1"},
					{Key: "class", Val: "test2"},
				},
			},
			key:  "id",
			want: "test1",
		},
		{
			name: "Test case 2",
			node: &html.Node{
				Type: html.ElementNode,
				Data: "div",
				Attr: []html.Attribute{
					{Key: "id", Val: "test1"},
					{Key: "class", Val: "test2"},
				},
			},
			key:  "class",
			want: "test2",
		},
		{
			name: "Test case 3",
			node: &html.Node{
				Type: html.ElementNode,
				Data: "div",
				Attr: []html.Attribute{
					{Key: "id", Val: "test1"},
					{Key: "class", Val: "test2"},
				},
			},
			key:  "name",
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

			if got := attr(tt.node, tt.key); got != tt.want {
				t.Errorf("attr() = %v, want %v", got, tt.want)
			}
		})
	}
}
