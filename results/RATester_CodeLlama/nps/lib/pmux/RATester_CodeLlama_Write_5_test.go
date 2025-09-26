package pmux

import (
	"fmt"
	"testing"
)

func TestWrite_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pConn := &PortConn{}
	b := []byte("test")
	n, err := pConn.Write(b)
	if err != nil {
		t.Errorf("Write failed, err: %v", err)
	}
	if n != len(b) {
		t.Errorf("Write failed, n: %d, want: %d", n, len(b))
	}
}
