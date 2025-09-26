package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestName_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := OutputFormat{
		Format: output.Format{
			Name: "name",
		},
	}
	if o.Name() != "name" {
		t.Errorf("Expected name to be %q but was %q", "name", o.Name())
	}
}
