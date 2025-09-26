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

	testDeps := &deps.Deps{}
	ns := New(testDeps)

	if ns.deps != testDeps {
		t.Errorf("Expected deps to be %v, got %v", testDeps, ns.deps)
	}
}
