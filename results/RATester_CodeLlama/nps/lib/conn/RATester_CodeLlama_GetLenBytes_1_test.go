package conn

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetLenBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := []byte{1, 2, 3, 4, 5}
	b, err := GetLenBytes(buf)
	if err != nil {
		t.Errorf("GetLenBytes failed, err: %v", err)
		return
	}
	if len(b) != 9 {
		t.Errorf("GetLenBytes failed, len(b): %d", len(b))
		return
	}
	if b[0] != 0 || b[1] != 0 || b[2] != 0 || b[3] != 5 {
		t.Errorf("GetLenBytes failed, b: %v", b)
		return
	}
	if !bytes.Equal(b[4:], buf) {
		t.Errorf("GetLenBytes failed, b: %v", b)
		return
	}
}
