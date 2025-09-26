package template

import (
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestArbitraryTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{
		ns: &nameSpace{
			set: map[string]*Template{
				"template1": {
					text: &template.Template{},
				},
				"template2": {
					text: &template.Template{},
				},
			},
		},
	}

	expected := e.ns.set["template1"]
	result := e.arbitraryTemplate()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
