package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/paerser/parser"
)

func TestEncodeNode_1(t *testing.T) {
	type args struct {
		labels map[string]string
		root   string
		node   *parser.Node
	}
	tests := []struct {
		name string
		args args
		want map[string]string
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

			encodeNode(tt.args.labels, tt.args.root, tt.args.node)
			if !reflect.DeepEqual(tt.args.labels, tt.want) {
				t.Errorf("encodeNode() = %v, want %v", tt.args.labels, tt.want)
			}
		})
	}
}
