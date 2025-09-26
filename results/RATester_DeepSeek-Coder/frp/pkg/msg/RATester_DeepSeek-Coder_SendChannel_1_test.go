package msg

import (
	"fmt"
	"testing"
)

func TestSendChannel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Dispatcher{
		sendCh: make(chan Message, 1),
	}

	expectedCh := d.sendCh
	actualCh := d.SendChannel()

	if expectedCh != actualCh {
		t.Errorf("Expected channel %v, but got %v", expectedCh, actualCh)
	}
}
