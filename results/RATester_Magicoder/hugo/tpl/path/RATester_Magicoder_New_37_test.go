package path

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNew_37(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	deps := &deps.Deps{
		// Initialize dependencies as needed
	}

	ns := New(deps)

	if ns == nil {
		t.Error("Expected a Namespace, but got nil")
	}

	if ns.deps != deps {
		t.Errorf("Expected deps to be %v, but got %v", deps, ns.deps)
	}
}
