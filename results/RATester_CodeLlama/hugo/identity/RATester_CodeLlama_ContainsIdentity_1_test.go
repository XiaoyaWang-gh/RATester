package identity

import (
	"fmt"
	"testing"
)

func TestContainsIdentity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	im := &identityManager{
		Identity: Anonymous,
	}

	r := im.ContainsIdentity(Anonymous)

	if r != FinderFound {
		t.Errorf("Expected FinderFound, got %v", r)
	}
}
