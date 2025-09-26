package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestGet_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := OutputFormats{
		{
			Format: output.Format{
				Name: "html",
			},
		},
		{
			Format: output.Format{
				Name: "json",
			},
		},
	}

	if o.Get("html") == nil {
		t.Error("Expected to find html format")
	}

	if o.Get("json") == nil {
		t.Error("Expected to find json format")
	}

	if o.Get("xml") != nil {
		t.Error("Expected to not find xml format")
	}
}
