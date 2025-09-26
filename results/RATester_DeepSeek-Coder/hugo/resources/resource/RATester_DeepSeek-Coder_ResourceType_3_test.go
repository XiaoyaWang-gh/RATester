package resource

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestResourceType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	holder := resourceTypesHolder{
		mediaType:    media.Type{},
		resourceType: "test",
	}

	result := holder.ResourceType()
	expected := "test"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
