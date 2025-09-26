package udp

import (
	"fmt"
	"testing"
)

func TestAccept_1(t *testing.T) {
	listener := &Listener{
		acceptCh: make(chan *Conn),
	}

	t.Run("should return a connection", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		conn := &Conn{}
		listener.acceptCh <- conn

		result, err := listener.Accept()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result != conn {
			t.Fatalf("expected %v, got %v", conn, result)
		}
	})

	t.Run("should return an error when listener is closed", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		close(listener.acceptCh)

		_, err := listener.Accept()
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
	})
}
