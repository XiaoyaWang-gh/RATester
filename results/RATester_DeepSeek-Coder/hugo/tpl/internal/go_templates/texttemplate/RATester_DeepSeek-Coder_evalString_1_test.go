package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEvalString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &state{
		node: &parse.StringNode{Text: "test"},
	}
	typ := reflect.TypeOf("")
	result := s.evalString(typ, s.node)
	expected := reflect.ValueOf("test")
	if result.String() != expected.String() {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
