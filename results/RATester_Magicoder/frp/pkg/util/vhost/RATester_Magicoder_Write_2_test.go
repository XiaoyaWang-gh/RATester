package vhost

import (
	"fmt"
	"io"
	"testing"
)

func TestWrite_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := readOnlyConn{}
	n, err := conn.Write([]byte("test"))
	if n != 0 || err != io.ErrClosedPipe {
		t.Errorf("Expected n=0 and error=io.ErrClosedPipe, but got n=%d and error=%v", n, err)
	}
}
