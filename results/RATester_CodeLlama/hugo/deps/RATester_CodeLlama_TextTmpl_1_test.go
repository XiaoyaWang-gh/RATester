package deps

import (
	"fmt"
	"testing"
)

func TestTextTmpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var d *Deps
	if d.TextTmpl() != d.tmplHandlers.TxtTmpl {
		t.Errorf("expected %v, got %v", d.tmplHandlers.TxtTmpl, d.TextTmpl())
	}
}
