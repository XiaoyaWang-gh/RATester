package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestAccept_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	httpsListener := &HttpsListener{
		acceptConn: make(chan *conn.Conn, 1),
	}

	testConn := &conn.Conn{}
	httpsListener.acceptConn <- testConn

	conn, err := httpsListener.Accept()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if conn != testConn {
		t.Fatalf("expected %v, got %v", testConn, conn)
	}
}
