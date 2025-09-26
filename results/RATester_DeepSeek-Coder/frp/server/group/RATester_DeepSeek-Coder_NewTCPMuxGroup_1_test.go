package group

import (
	"fmt"
	"testing"
)

func TestNewTCPMuxGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &TCPMuxGroupCtl{
		groups: make(map[string]*TCPMuxGroup),
	}

	muxGroup := NewTCPMuxGroup(ctl)

	if muxGroup.ctl != ctl {
		t.Errorf("Expected ctl to be %v, got %v", ctl, muxGroup.ctl)
	}

	if len(muxGroup.lns) != 0 {
		t.Errorf("Expected lns to be empty, got %v", muxGroup.lns)
	}

	if muxGroup.acceptCh == nil {
		t.Errorf("Expected acceptCh to be initialized")
	}
}
