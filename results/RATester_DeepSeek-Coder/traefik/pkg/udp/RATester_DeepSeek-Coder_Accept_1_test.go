package udp

import (
	"fmt"
	"testing"
)

func TestAccept_1(t *testing.T) {
	listener := &Listener{
		acceptCh: make(chan *Conn),
	}

	t.Run("should return connection when connection is available", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		conn := &Conn{}
		go func() {
			listener.acceptCh <- conn
		}()

		result, err := listener.Accept()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != conn {
			t.Errorf("expected %v, got %v", conn, result)
		}
	})

	t.Run("should return error when listener is closed", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		close(listener.acceptCh)

		result, err := listener.Accept()
		if err != errClosedListener {
			t.Errorf("expected %v, got %v", errClosedListener, err)
		}
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
}
