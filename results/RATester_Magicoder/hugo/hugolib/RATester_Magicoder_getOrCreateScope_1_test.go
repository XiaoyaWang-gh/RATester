package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/output"
)

func TestgetOrCreateScope_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scope := "testScope"
	pco := &pageContentOutput{
		po: &pageOutput{
			f: output.Format{},
		},
	}
	c := &cachedContent{
		scopes: maps.NewCache[string, *cachedContentScope](),
	}

	cs := c.getOrCreateScope(scope, pco)

	if cs == nil {
		t.Error("Expected a cachedContentScope, but got nil")
	}

	if cs.scope != scope {
		t.Errorf("Expected scope to be %s, but got %s", scope, cs.scope)
	}

	if cs.pco != pco {
		t.Errorf("Expected pco to be %v, but got %v", pco, cs.pco)
	}

	if cs.cachedContent != c {
		t.Errorf("Expected cachedContent to be %v, but got %v", c, cs.cachedContent)
	}
}
