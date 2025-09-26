package kv

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/paerser/parser"
)

func TestContainsNode_1(t *testing.T) {
	type args struct {
		nodes []*parser.Node
		name  string
	}
	tests := []struct {
		name string
		args args
		want *parser.Node
	}{
		{
			name: "Test Case 1",
			args: args{
				nodes: []*parser.Node{
					{Name: "Node1"},
					{Name: "Node2"},
					{Name: "Node3"},
				},
				name: "Node2",
			},
			want: &parser.Node{Name: "Node2"},
		},
		{
			name: "Test Case 2",
			args: args{
				nodes: []*parser.Node{
					{Name: "Node1"},
					{Name: "Node2"},
					{Name: "Node3"},
				},
				name: "Node4",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := containsNode(tt.args.nodes, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("containsNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
