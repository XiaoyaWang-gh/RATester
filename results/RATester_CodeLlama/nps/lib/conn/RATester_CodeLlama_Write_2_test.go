package conn

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWrite_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	conn := &LenConn{
		conn: &bytes.Buffer{},
	}
	n, err := conn.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatalf("n = %d, want 5", n)
	}
	if conn.Len != 5 {
		t.Fatalf("conn.Len = %d, want 5", conn.Len)
	}
}
