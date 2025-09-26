package resources

import (
	"fmt"
	"testing"
)

func TestSetName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &metaResource{
		name:    "oldName",
		changed: false,
	}

	r.setName("newName")

	if r.name != "newName" {
		t.Errorf("Expected name to be 'newName', got '%s'", r.name)
	}

	if !r.changed {
		t.Errorf("Expected changed to be true, got false")
	}
}
