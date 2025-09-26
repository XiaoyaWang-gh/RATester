package group

import (
	"fmt"
	"net"
	"testing"
)

func TestCloseListener_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	testGroup := &TCPMuxGroup{
		group:    "testGroup",
		lns:      make([]*TCPMuxGroupListener, 0),
		acceptCh: make(chan net.Conn),
	}

	testListener := &TCPMuxGroupListener{
		groupName: "testGroup",
		group:     testGroup,
		closeCh:   make(chan struct{}),
	}

	testGroup.lns = append(testGroup.lns, testListener)

	testGroup.CloseListener(testListener)

	if len(testGroup.lns) != 0 {
		t.Errorf("Expected lns length to be 0 after CloseListener, got %d", len(testGroup.lns))
	}

	_, ok := <-testGroup.acceptCh
	if ok {
		t.Errorf("Expected acceptCh to be closed after CloseListener")
	}
}
