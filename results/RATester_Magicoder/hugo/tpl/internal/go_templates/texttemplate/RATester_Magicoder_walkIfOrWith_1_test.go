package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestwalkIfOrWith_1(t *testing.T) {
	type args struct {
		typ      parse.NodeType
		dot      reflect.Value
		pipe     *parse.PipeNode
		list     *parse.ListNode
		elseList *parse.ListNode
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

			s := &state{}
			s.walkIfOrWith(tt.args.typ, tt.args.dot, tt.args.pipe, tt.args.list, tt.args.elseList)
		})
	}
}
