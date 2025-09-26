package templates

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNew_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	deps := &deps.Deps{}
	ns := New(deps)

	if ns.deps != deps {
		t.Errorf("Expected deps to be %v, but got %v", deps, ns.deps)
	}
}
