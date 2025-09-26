package kv

import (
	"fmt"
	"testing"

	"github.com/traefik/paerser/parser"
)

func TestDecodeToNode_1(t *testing.T) {
	type args struct {
		root  *parser.Node
		path  []string
		value string
	}
	tests := []struct {
		name string
		args args
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

			decodeToNode(tt.args.root, tt.args.path, tt.args.value)
		})
	}
}
