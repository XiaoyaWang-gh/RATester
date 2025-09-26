package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestWalkIfOrWith_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var s *state
	var typ parse.NodeType
	var dot reflect.Value
	var pipe *parse.PipeNode
	var list *parse.ListNode
	var elseList *parse.ListNode
	s = &state{}
	typ = parse.NodeIf
	dot = reflect.ValueOf(nil)
	pipe = &parse.PipeNode{}
	list = &parse.ListNode{}
	elseList = &parse.ListNode{}
	s.walkIfOrWith(typ, dot, pipe, list, elseList)
}
