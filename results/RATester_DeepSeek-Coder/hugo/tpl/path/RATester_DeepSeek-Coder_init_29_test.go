package path

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestInit_29(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	ctx := New(d)
	if ctx == nil {
		t.Errorf("Expected a non-nil context")
	}
}
