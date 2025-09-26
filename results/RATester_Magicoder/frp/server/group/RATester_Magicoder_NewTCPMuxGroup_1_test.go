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

	if muxGroup == nil {
		t.Error("NewTCPMuxGroup returned nil")
	}

	if muxGroup.ctl != ctl {
		t.Error("NewTCPMuxGroup did not set the correct ctl")
	}

	if len(muxGroup.lns) != 0 {
		t.Error("NewTCPMuxGroup did not initialize lns correctly")
	}

	if muxGroup.acceptCh == nil {
		t.Error("NewTCPMuxGroup did not initialize acceptCh correctly")
	}
}
