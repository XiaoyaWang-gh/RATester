package template

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestarbitraryTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{
		ns: &nameSpace{
			set: map[string]*Template{
				"template1": {
					Tree: &parse.Tree{},
				},
				"template2": {
					Tree: &parse.Tree{},
				},
			},
		},
	}

	expected := e.ns.set["template1"]
	actual := e.arbitraryTemplate()

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
