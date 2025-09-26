package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func Testinit_33(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	ctx, err := New(d)
	if err != nil {
		t.Fatal(err)
	}

	if ctx.cssNs == nil {
		t.Error("css namespace not found")
	}

	if ctx.jsNs == nil {
		t.Error("js namespace not found")
	}
}
