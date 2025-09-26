package udp

import (
	"fmt"
	"testing"
)

func TestWrite_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Conn{}
	p := []byte{}
	n, err := c.Write(p)
	if n != 0 {
		t.Errorf("n = %v, want 0", n)
	}
	if err != nil {
		t.Errorf("err = %v, want nil", err)
	}
}
