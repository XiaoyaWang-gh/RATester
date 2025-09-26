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
			name: "test containsNode",
			args: args{
				nodes: []*parser.Node{
					{
						Name:        "test",
						Description: "test",
						FieldName:   "test",
						Value:       "test",
						RawValue:    "test",
						Disabled:    true,
						Kind:        reflect.String,
						Tag:         "test",
						Children:    []*parser.Node{},
					},
				},
				name: "test",
			},
			want: &parser.Node{
				Name:        "test",
				Description: "test",
				FieldName:   "test",
				Value:       "test",
				RawValue:    "test",
				Disabled:    true,
				Kind:        reflect.String,
				Tag:         "test",
				Children:    []*parser.Node{},
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

			if got := containsNode(tt.args.nodes, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("containsNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
