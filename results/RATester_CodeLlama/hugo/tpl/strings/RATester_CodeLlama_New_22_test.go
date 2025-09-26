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
	n := New(d)
	if n == nil {
		t.Fatal("New returned nil")
	}
	if n.deps != d {
		t.Errorf("New returned %v, expected %v", n.deps, d)
	}
}
