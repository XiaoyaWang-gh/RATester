package conn

import (
	"fmt"
	"testing"
)

func TestRead_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SnappyConn{}
	b := []byte{}
	n, err := s.Read(b)
	if err != nil {
		t.Errorf("Read() error = %v", err)
		return
	}
	if n != 0 {
		t.Errorf("Read() n = %v, want 0", n)
	}
}
