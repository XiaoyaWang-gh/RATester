package strings

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNew_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	ns := New(d)
	if ns.deps != d {
		t.Errorf("Expected deps to be %v, got %v", d, ns.deps)
	}
}
