package deps

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestSetTempl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	d := &Deps{
		tmplHandlers: &tpl.TemplateHandlers{},
	}

	newTmplHandlers := &tpl.TemplateHandlers{}
	d.SetTempl(newTmplHandlers)

	if d.tmplHandlers != newTmplHandlers {
		t.Errorf("Expected tmplHandlers to be %v, but got %v", newTmplHandlers, d.tmplHandlers)
	}
}
