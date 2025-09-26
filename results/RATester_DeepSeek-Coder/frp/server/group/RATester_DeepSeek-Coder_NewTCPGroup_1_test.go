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
		t.Errorf("NewTCPGroup() = %v, want not nil", group)
	}

	if group.ctl != ctl {
		t.Errorf("NewTCPGroup().ctl = %v, want %v", group.ctl, ctl)
	}

	if len(group.lns) != 0 {
		t.Errorf("NewTCPGroup().lns = %v, want empty", group.lns)
	}

	if group.acceptCh == nil {
		t.Errorf("NewTCPGroup().acceptCh = %v, want not nil", group.acceptCh)
	}
}
