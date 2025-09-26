package template

import (
	"fmt"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func Testcommit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &escaper{
		output:            make(map[string]context),
		derived:           make(map[string]*template.Template),
		called:            make(map[string]bool),
		actionNodeEdits:   make(map[*parse.ActionNode][]string),
		templateNodeEdits: make(map[*parse.TemplateNode]string),
		textNodeEdits:     make(map[*parse.TextNode][]byte),
	}

	// Add some test data to the escaper
	e.output["testTemplate"] = context{}
	e.derived["testDerived"] = &template.Template{}
	e.called["testCalled"] = true
	e.actionNodeEdits[&parse.ActionNode{}] = []string{"testAction"}
	e.templateNodeEdits[&parse.TemplateNode{}] = "testTemplateNode"
	e.textNodeEdits[&parse.TextNode{}] = []byte("testTextNode")

	// Call the method under test
	e.commit()

	// Assert that the escaper has been reset as expected
	if len(e.output) != 0 {
		t.Errorf("Expected output map to be empty, but it was not")
	}
	if len(e.derived) != 0 {
		t.Errorf("Expected derived map to be empty, but it was not")
	}
	if len(e.called) != 0 {
		t.Errorf("Expected called map to be empty, but it was not")
	}
	if len(e.actionNodeEdits) != 0 {
		t.Errorf("Expected actionNodeEdits map to be empty, but it was not")
	}
	if len(e.templateNodeEdits) != 0 {
		t.Errorf("Expected templateNodeEdits map to be empty, but it was not")
	}
	if len(e.textNodeEdits) != 0 {
		t.Errorf("Expected textNodeEdits map to be empty, but it was not")
	}
}
