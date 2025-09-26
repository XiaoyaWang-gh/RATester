package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestClone_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	te := templateExec{}
	cloned := te.Clone(d)

	if cloned.d != d {
		t.Errorf("Expected cloned templateExec to have the same deps as the original")
	}

	if cloned.executor != nil {
		t.Errorf("Expected cloned templateExec to have nil executor")
	}

	if cloned.funcs != nil {
		t.Errorf("Expected cloned templateExec to have nil funcs")
	}

	if cloned.templateHandler != nil {
		t.Errorf("Expected cloned templateExec to have nil templateHandler")
	}
}
