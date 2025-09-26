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

	im := &identityManager{
		name: "test",
	}

	manager := im.GetDependencyManagerForScope(0)

	if manager != im {
		t.Errorf("Expected manager to be %v, got %v", im, manager)
	}
}
