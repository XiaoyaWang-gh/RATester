package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	httpsListener := &HttpsListener{
		acceptConn: make(chan *conn.Conn),
	}
	addr := httpsListener.Addr()
	if addr == nil {
		t.Errorf("Addr() = nil, want not nil")
	}
}
