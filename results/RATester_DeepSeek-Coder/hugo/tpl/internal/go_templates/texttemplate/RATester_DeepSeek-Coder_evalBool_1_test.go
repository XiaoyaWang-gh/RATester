package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &state{
		node: &parse.BoolNode{
			True: true,
		},
	}
	typ := reflect.TypeOf(true)
	result := s.evalBool(typ, s.node)
	expected := reflect.ValueOf(true)
	if result.Bool() != expected.Bool() {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
