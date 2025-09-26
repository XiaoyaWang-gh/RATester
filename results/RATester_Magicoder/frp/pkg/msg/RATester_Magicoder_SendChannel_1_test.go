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
		sendCh: make(chan Message),
	}

	expected := d.sendCh
	actual := d.SendChannel()

	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
