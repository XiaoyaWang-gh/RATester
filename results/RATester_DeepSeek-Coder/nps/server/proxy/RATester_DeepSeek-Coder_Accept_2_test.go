package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestAccept_2(t *testing.T) {
	httpsListener := &HttpsListener{
		acceptConn: make(chan *conn.Conn, 1),
	}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		mockConn := &conn.Conn{}
		httpsListener.acceptConn <- mockConn
		conn, err := httpsListener.Accept()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if conn != mockConn {
			t.Errorf("expected %v, got %v", mockConn, conn)
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		close(httpsListener.acceptConn)
		conn, err := httpsListener.Accept()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if conn != nil {
			t.Errorf("expected nil, got %v", conn)
		}
	})
}
