package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
	}
	data := []byte("Hello, World!")
	n, err := w.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(data) {
		t.Errorf("wrote %d bytes; want %d", n, len(data))
	}
	if w.size != n {
		t.Errorf("size %d; want %d", w.size, n)
	}
	if w.status != 200 {
		t.Errorf("status %d; want 200", w.status)
	}
}
