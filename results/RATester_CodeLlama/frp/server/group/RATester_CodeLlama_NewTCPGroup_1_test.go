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
	tg := NewTCPGroup(ctl)
	if tg == nil {
		t.Error("NewTCPGroup failed")
	}
}
