package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestsetName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &metaResource{
		changed: false,
		title:   "Test Title",
		name:    "Test Name",
		params:  maps.Params{},
	}

	r.setName("New Name")

	if r.name != "New Name" {
		t.Errorf("Expected name to be 'New Name', but got '%s'", r.name)
	}

	if !r.changed {
		t.Error("Expected changed to be true")
	}
}
