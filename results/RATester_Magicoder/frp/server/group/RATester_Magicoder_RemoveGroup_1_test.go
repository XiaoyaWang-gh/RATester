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

	// Test the method
	tmgc.RemoveGroup("testGroup")

	// Check if the group is removed
	if _, ok := tmgc.groups["testGroup"]; ok {
		t.Error("Group was not removed")
	}
}
