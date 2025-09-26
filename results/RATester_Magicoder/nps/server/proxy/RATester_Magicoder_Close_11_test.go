package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestClose_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	httpsListener := &HttpsListener{
		acceptConn: make(chan *conn.Conn),
	}

	err := httpsListener.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
