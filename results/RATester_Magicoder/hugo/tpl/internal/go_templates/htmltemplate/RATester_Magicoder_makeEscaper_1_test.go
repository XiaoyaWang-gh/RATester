package template

import (
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestmakeEscaper_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &nameSpace{
		set: map[string]*Template{
			"test": {
				text: &template.Template{},
			},
		},
	}
	esc := makeEscaper(ns)

	if esc.ns != ns {
		t.Errorf("Expected ns to be %v, but got %v", ns, esc.ns)
	}

	if len(esc.output) != 0 {
		t.Errorf("Expected output to be empty, but got %v", esc.output)
	}

	if len(esc.derived) != 0 {
		t.Errorf("Expected derived to be empty, but got %v", esc.derived)
	}

	if len(esc.called) != 0 {
		t.Errorf("Expected called to be empty, but got %v", esc.called)
	}

	if len(esc.actionNodeEdits) != 0 {
		t.Errorf("Expected actionNodeEdits to be empty, but got %v", esc.actionNodeEdits)
	}

	if len(esc.templateNodeEdits) != 0 {
		t.Errorf("Expected templateNodeEdits to be empty, but got %v", esc.templateNodeEdits)
	}

	if len(esc.textNodeEdits) != 0 {
		t.Errorf("Expected textNodeEdits to be empty, but got %v", esc.textNodeEdits)
	}

	if esc.rangeContext != nil {
		t.Errorf("Expected rangeContext to be nil, but got %v", esc.rangeContext)
	}
}
