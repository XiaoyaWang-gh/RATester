package group

import (
	"fmt"
	"testing"
)

func TestCloseListener_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tmg := &TCPMuxGroup{
		group: "testGroup",
		lns: []*TCPMuxGroupListener{
			{groupName: "testGroup1"},
			{groupName: "testGroup2"},
		},
	}

	ln := &TCPMuxGroupListener{groupName: "testGroup1"}
	tmg.CloseListener(ln)

	if len(tmg.lns) != 1 {
		t.Errorf("Expected 1 listener, got %d", len(tmg.lns))
	}

	if tmg.lns[0].groupName != "testGroup2" {
		t.Errorf("Expected 'testGroup2', got '%s'", tmg.lns[0].groupName)
	}
}
