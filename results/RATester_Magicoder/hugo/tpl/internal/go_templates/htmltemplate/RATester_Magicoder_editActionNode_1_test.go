package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TesteditActionNode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{
		actionNodeEdits: make(map[*parse.ActionNode][]string),
	}
	n := &parse.ActionNode{
		Pipe: &parse.PipeNode{},
	}
	cmds := []string{"cmd1", "cmd2"}
	e.editActionNode(n, cmds)
	if _, ok := e.actionNodeEdits[n]; !ok {
		t.Errorf("expected actionNodeEdits to contain node %v", n)
	}
	if !reflect.DeepEqual(e.actionNodeEdits[n], cmds) {
		t.Errorf("expected actionNodeEdits[%v] to be %v, got %v", n, cmds, e.actionNodeEdits[n])
	}
}
