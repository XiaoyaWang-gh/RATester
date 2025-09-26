package images

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func Testinit_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	ctx := New(d)

	if ctx == nil {
		t.Error("Expected a non-nil context, but got nil")
	}
}
