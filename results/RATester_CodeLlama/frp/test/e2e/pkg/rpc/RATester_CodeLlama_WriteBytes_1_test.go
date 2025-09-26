package rpc

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriteBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := []byte("hello")
	w := &bytes.Buffer{}
	n, err := WriteBytes(w, buf)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(buf) {
		t.Fatalf("expected %d, got %d", len(buf), n)
	}
	if w.Len() != len(buf) {
		t.Fatalf("expected %d, got %d", len(buf), w.Len())
	}
	if !bytes.Equal(w.Bytes(), buf) {
		t.Fatalf("expected %v, got %v", buf, w.Bytes())
	}
}
