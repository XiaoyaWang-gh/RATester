package pmux

import (
	"fmt"
	"net"
	"testing"
)

func TestAccept_3(t *testing.T) {
	listener := &PortListener{
		connCh:  make(chan *PortConn, 1),
		isClose: false,
	}

	t.Run("TestAccept_WhenListenerIsClosed", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		listener.isClose = true
		_, err := listener.Accept()
		if err == nil {
			t.Error("Expected error, but got nil")
		}
		if err.Error() != "the listener has closed" {
			t.Errorf("Expected error message 'the listener has closed', but got '%s'", err.Error())
		}
	})

	t.Run("TestAccept_WhenListenerHasConn", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		listener.isClose = false
		conn := &PortConn{
			Conn:     &net.TCPConn{},
			rs:       []byte("test"),
			readMore: false,
			start:    0,
		}
		listener.connCh <- conn
		_, err := listener.Accept()
		if err != nil {
			t.Errorf("Expected no error, but got '%s'", err.Error())
		}
	})

	t.Run("TestAccept_WhenListenerHasNoConn", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		listener.isClose = false
		listener.connCh <- nil
		_, err := listener.Accept()
		if err == nil {
			t.Error("Expected error, but got nil")
		}
		if err.Error() != "the listener has closed" {
			t.Errorf("Expected error message 'the listener has closed', but got '%s'", err.Error())
		}
	})
}
