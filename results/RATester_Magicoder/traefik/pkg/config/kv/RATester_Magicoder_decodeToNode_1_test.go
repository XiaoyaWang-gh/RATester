package kv

import (
	"fmt"
	"reflect"
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
		want *parser.Node
	}{
		{
			name: "Test case 1",
			args: args{
				root:  &parser.Node{},
				path:  []string{"field1", "subfield1"},
				value: "value1",
			},
			want: &parser.Node{
				Name:  "field1",
				Value: "",
				Children: []*parser.Node{
					{
						Name:     "subfield1",
						Value:    "value1",
						Children: []*parser.Node{},
					},
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				root:  &parser.Node{},
				path:  []string{"field2", "subfield2"},
				value: "value2",
			},
			want: &parser.Node{
				Name:  "field2",
				Value: "",
				Children: []*parser.Node{
					{
						Name:     "subfield2",
						Value:    "value2",
						Children: []*parser.Node{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			decodeToNode(tt.args.root, tt.args.path, tt.args.value)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("decodeToNode() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
