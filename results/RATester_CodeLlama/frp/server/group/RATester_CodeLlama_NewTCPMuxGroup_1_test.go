package group

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
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
	tmg := NewTCPMuxGroup(ctl)
	assert.NotNil(t, tmg)
}
