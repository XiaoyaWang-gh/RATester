package internal

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func Testattr_1(t *testing.T) {
	node := &html.Node{
		Attr: []html.Attribute{
			{Key: "key1", Val: "value1"},
			{Key: "key2", Val: "value2"},
		},
	}

	tests := []struct {
		name string
		node *html.Node
		key  string
		want string
	}{
		{
			name: "Test case 1",
			node: node,
			key:  "key1",
			want: "value1",
		},
		{
			name: "Test case 2",
			node: node,
			key:  "key2",
			want: "value2",
		},
		{
			name: "Test case 3",
			node: node,
			key:  "key3",
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
