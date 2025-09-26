package livereload

import (
	"fmt"
	"testing"

	"github.com/gorilla/websocket"
)

func TestReader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connection{}
	c.ws = &websocket.Conn{}
	c.send = make(chan []byte)
	c.reader()
	// TODO: add assertions
}
