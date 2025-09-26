package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestsetTitle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &metaResource{
		changed: false,
		title:   "",
		name:    "",
		params:  maps.Params{},
	}

	r.setTitle("Test Title")

	if r.title != "Test Title" {
		t.Errorf("Expected title to be 'Test Title', but got '%s'", r.title)
	}

	if !r.changed {
		t.Error("Expected changed to be true after setting title")
	}
}
