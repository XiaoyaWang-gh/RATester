package fiber

import (
	"fmt"
	"testing"
)

func TestWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &testConn{}
	b := []byte("hello")
	n, err := c.Write(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(b) {
		t.Fatalf("wrote %d bytes, expected %d", n, len(b))
	}
	if c.w.String() != "hello" {
		t.Fatalf("wrote %q, expected %q", c.w.String(), "hello")
	}
}
