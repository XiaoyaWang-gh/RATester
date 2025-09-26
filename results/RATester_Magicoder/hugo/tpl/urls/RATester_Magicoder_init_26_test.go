package urls

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func Testinit_26(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	ctx := New(d)

	if ctx == nil {
		t.Error("Expected ctx to be not nil")
	}
}
