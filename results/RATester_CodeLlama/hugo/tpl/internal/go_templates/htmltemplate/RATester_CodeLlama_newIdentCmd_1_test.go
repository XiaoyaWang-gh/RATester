package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestNewIdentCmd_1(t *testing.T) {
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
			name: "test1",
			args: args{
				identifier: "test1",
				pos:        parse.Pos(0),
			},
			want: &parse.CommandNode{
				NodeType: parse.NodeCommand,
				Args:     []parse.Node{parse.NewIdentifier("test1").SetTree(nil).SetPos(parse.Pos(0))},
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
