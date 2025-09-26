package tplimpl

import (
	"fmt"
	"testing"

	htmltemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/htmltemplate"
	texttemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestcreatePrototypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &templateNamespace{
		prototypeText:      texttemplate.Must(texttemplate.New("test").Parse("Hello, {{.}}")),
		prototypeHTML:      htmltemplate.Must(htmltemplate.New("test").Parse("Hello, {{.}}")),
		prototypeTextClone: nil,
		prototypeHTMLClone: nil,
	}

	err := ns.createPrototypes()
	if err != nil {
		t.Errorf("createPrototypes() returned an error: %v", err)
	}

	if ns.prototypeTextClone == nil || ns.prototypeHTMLClone == nil {
		t.Errorf("createPrototypes() did not clone the templates correctly")
	}
}
