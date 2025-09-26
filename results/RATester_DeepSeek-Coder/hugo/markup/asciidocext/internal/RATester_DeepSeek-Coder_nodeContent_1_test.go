package internal

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestNodeContent_1(t *testing.T) {
	type args struct {
		node *html.Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := nodeContent(tt.args.node); got != tt.want {
				t.Errorf("nodeContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
