package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestReadFlag_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: &net.TCPConn{},
		Rb:   []byte{},
	}

	flag, err := conn.ReadFlag()
	if err != nil {
		t.Errorf("ReadFlag() error = %v, wantErr %v", err, nil)
		return
	}
	if flag != "" {
		t.Errorf("ReadFlag() = %v, want %v", flag, "")
	}
}
