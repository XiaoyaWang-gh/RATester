package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestPermalink_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := OutputFormat{
		Rel: "canonical",
		Format: output.Format{
			Name: "test",
		},
		relPermalink: "/test",
		permalink:    "/test",
	}

	expected := "/test"
	result := o.Permalink()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
