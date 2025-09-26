package fiber

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tc := &testConn{}
	data := []byte("Hello, World!")
	n, err := tc.Write(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != len(data) {
		t.Errorf("Expected to write %d bytes, but wrote %d", len(data), n)
	}
	if !bytes.Equal(tc.w.Bytes(), data) {
		t.Errorf("Expected to write %v, but wrote %v", data, tc.w.Bytes())
	}
}
