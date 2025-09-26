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
	httpsConn := &conn.Conn{}
	httpsListener.acceptConn <- httpsConn
	conn, err := httpsListener.Accept()
	if err != nil {
		t.Error(err)
	}
	if conn != httpsConn {
		t.Error("conn != httpsConn")
	}
}
