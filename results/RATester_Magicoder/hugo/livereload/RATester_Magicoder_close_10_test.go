package livereload

import (
	"fmt"
	"testing"
)

func Testclose_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &connection{
		send: make(chan []byte, 1),
	}

	conn.close()

	_, ok := <-conn.send
	if ok {
		t.Error("Expected channel to be closed")
	}
}
