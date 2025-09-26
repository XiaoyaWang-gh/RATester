package os

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNew_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	n := New(d)

	if n.readFileFs == nil {
		t.Errorf("readFileFs should not be nil")
	}

	if n.workFs == nil {
		t.Errorf("workFs should not be nil")
	}

	if n.deps == nil {
		t.Errorf("deps should not be nil")
	}
}
