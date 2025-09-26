package group

import (
	"fmt"
	"testing"
)

func TestRemoveGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tmgc := &TCPMuxGroupCtl{
		groups: make(map[string]*TCPMuxGroup),
	}

	// Add a group
	tmgc.groups["testGroup"] = &TCPMuxGroup{}

	// Remove the group
	tmgc.RemoveGroup("testGroup")

	// Check if the group is removed
	_, exists := tmgc.groups["testGroup"]
	if exists {
		t.Errorf("Expected group to be removed, but it still exists")
	}
}
