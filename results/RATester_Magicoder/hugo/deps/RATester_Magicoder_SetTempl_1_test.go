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

	d := &Deps{}
	tmplHandlers := &tpl.TemplateHandlers{}

	d.SetTempl(tmplHandlers)

	if d.tmplHandlers != tmplHandlers {
		t.Errorf("Expected tmplHandlers to be set, but it was not.")
	}
}
