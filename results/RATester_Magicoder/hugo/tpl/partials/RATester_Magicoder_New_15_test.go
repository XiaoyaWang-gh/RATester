package partials

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNew_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	deps := &deps.Deps{}
	ns := New(deps)

	if ns == nil {
		t.Error("Expected a Namespace, got nil")
	}

	if ns.deps != deps {
		t.Errorf("Expected deps to be %v, got %v", deps, ns.deps)
	}

	if ns.cachedPartials == nil {
		t.Error("Expected a non-nil cachedPartials")
	}
}
