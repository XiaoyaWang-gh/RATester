package deps

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestTmpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Deps{
		tmplHandlers: &tpl.TemplateHandlers{},
	}

	expected := d.tmplHandlers.Tmpl
	actual := d.Tmpl()

	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
