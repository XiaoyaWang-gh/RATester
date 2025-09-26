package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestnewIdentCmd_1(t *testing.T) {
	type args struct {
		identifier string
		pos        parse.Pos
	}
	tests := []struct {
		name string
		args args
		want *parse.CommandNode
	}{
		{
			name: "Test case 1",
			args: args{
				identifier: "test_identifier",
				pos:        1,
			},
			want: &parse.CommandNode{
				NodeType: parse.NodeCommand,
				Args:     []parse.Node{parse.NewIdentifier("test_identifier").SetTree(nil).SetPos(1)},
			},
		},
		{
			name: "Test case 2",
			args: args{
				identifier: "another_identifier",
				pos:        2,
			},
			want: &parse.CommandNode{
				NodeType: parse.NodeCommand,
				Args:     []parse.Node{parse.NewIdentifier("another_identifier").SetTree(nil).SetPos(2)},
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

			if got := newIdentCmd(tt.args.identifier, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newIdentCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
