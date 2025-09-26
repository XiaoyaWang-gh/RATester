package udp

import (
	"fmt"
	"testing"
	"time"
)

func TestReadLoop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// setup
	c := &Conn{}
	c.receiveCh = make(chan []byte, 1)
	c.readCh = make(chan []byte, 1)
	c.sizeCh = make(chan int, 1)
	c.msgs = make([][]byte, 0)
	c.timeout = time.Second
	c.doneCh = make(chan struct{})

	// test
	c.readLoop()

	// verify
	// ...
}
