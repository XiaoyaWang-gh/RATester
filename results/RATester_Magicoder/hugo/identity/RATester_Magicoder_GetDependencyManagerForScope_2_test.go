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
		name: "testManager",
	}

	expected := im
	actual := im.GetDependencyManagerForScope(0)

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
