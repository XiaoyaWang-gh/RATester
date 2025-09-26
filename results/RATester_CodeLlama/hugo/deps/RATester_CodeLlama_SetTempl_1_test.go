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
	tpl := &tpl.TemplateHandlers{}
	d.SetTempl(tpl)
	if d.tmplHandlers != tpl {
		t.Errorf("expected %v, got %v", tpl, d.tmplHandlers)
	}
}
