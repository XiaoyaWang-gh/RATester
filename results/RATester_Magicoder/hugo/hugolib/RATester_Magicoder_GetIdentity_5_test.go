package hugolib

import (
	"fmt"
	"testing"
)

func TestGetIdentity_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pageState := &pageState{
		pid: 1,
	}

	expected := pageState
	actual := pageState.GetIdentity()

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
