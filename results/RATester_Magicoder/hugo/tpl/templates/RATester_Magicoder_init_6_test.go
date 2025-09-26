package templates

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func Testinit_6(t *testing.T) {
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

	if ctx.Exists("partials/header.html") != true {
		t.Error("Expected Exists to return true for partials/header.html")
	}

	if ctx.Exists("partials/doesnotexist.html") != false {
		t.Error("Expected Exists to return false for partials/doesnotexist.html")
	}
}
