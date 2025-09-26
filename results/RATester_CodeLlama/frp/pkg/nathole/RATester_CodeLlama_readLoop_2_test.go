package nathole

import (
	"fmt"
	"testing"
)

func TestReadLoop_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	conn := &discoverConn{}
	conn.messageChan = make(chan *Message)
	// when
	conn.readLoop()
	// then
	// TODO
}
