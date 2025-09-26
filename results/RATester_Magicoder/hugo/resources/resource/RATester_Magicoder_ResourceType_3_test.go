package resource

import (
	"fmt"
	"testing"
)

func TestResourceType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	holder := resourceTypesHolder{
		resourceType: "test",
	}

	if holder.ResourceType() != "test" {
		t.Errorf("Expected 'test', got '%s'", holder.ResourceType())
	}
}
