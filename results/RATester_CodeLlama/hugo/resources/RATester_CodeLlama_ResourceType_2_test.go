package resources

import (
	"fmt"
	"testing"
)

func TestResourceType_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &genericResource{}
	if l.ResourceType() != l.MediaType().MainType {
		t.Errorf("ResourceType() = %v, want %v", l.ResourceType(), l.MediaType().MainType)
	}
}
