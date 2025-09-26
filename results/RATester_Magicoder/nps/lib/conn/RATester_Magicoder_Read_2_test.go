package conn

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/golang/snappy"
)

func TestRead_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &SnappyConn{
		r: snappy.NewReader(bytes.NewBuffer([]byte("test"))),
	}

	b := make([]byte, 4)
	n, err := conn.Read(b)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if n != 4 {
		t.Errorf("Expected to read 4 bytes, but got %d", n)
	}

	if string(b) != "test" {
		t.Errorf("Expected to read 'test', but got '%s'", string(b))
	}
}
