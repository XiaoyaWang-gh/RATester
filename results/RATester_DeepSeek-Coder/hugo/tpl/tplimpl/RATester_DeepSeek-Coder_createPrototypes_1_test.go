package tplimpl

import (
	"fmt"
	"testing"

	htmltemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/htmltemplate"
	texttemplate "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
)

func TestCreatePrototypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &templateNamespace{
		prototypeText: texttemplate.Must(texttemplate.New("test").Parse("test")),
		prototypeHTML: htmltemplate.Must(htmltemplate.New("test").Parse("test")),
	}

	err := ns.createPrototypes()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if ns.prototypeTextClone == nil || ns.prototypeHTMLClone == nil {
		t.Errorf("Expected prototypeTextClone and prototypeHTMLClone to be non-nil")
	}
}
