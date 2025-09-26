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
	clone := te.Clone(d)

	if clone.d != d {
		t.Errorf("Expected clone.d to be %v, but got %v", d, clone.d)
	}

	if clone.executor != nil {
		t.Errorf("Expected clone.executor to be nil, but got %v", clone.executor)
	}

	if clone.funcs != nil {
		t.Errorf("Expected clone.funcs to be nil, but got %v", clone.funcs)
	}
}
