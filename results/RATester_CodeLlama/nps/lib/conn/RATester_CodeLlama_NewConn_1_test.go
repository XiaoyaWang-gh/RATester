package conn

import (
	"fmt"
	"testing"
)

func TestNewConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{}
	newConn := NewConn(conn)
	if newConn.Conn != conn {
		t.Errorf("NewConn() = %v, want %v", newConn.Conn, conn)
	}
}
