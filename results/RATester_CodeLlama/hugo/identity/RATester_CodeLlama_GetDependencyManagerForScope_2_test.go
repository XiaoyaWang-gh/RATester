package identity

import (
	"fmt"
	"testing"
)

func TestGetDependencyManagerForScope_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	im := &identityManager{}
	if im.GetDependencyManagerForScope(0) != im {
		t.Errorf("Expected %v, got %v", im, im.GetDependencyManagerForScope(0))
	}
}
