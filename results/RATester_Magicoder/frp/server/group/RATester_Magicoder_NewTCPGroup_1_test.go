package group

import (
	"fmt"
	"testing"
)

func TestNewTCPGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &TCPGroupCtl{
		groups: make(map[string]*TCPGroup),
	}

	group := NewTCPGroup(ctl)

	if group == nil {
		t.Error("NewTCPGroup failed")
	}

	if group.ctl != ctl {
		t.Error("NewTCPGroup failed, ctl not set correctly")
	}

	if len(group.lns) != 0 {
		t.Error("NewTCPGroup failed, lns not initialized correctly")
	}

	if group.acceptCh == nil {
		t.Error("NewTCPGroup failed, acceptCh not initialized correctly")
	}
}
