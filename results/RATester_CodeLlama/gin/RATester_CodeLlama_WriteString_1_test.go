package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
	}
	n, err := w.WriteString("hello")
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Errorf("wrote %d bytes; expected 5", n)
	}
	if w.size != 5 {
		t.Errorf("w.size = %d; expected 5", w.size)
	}
	if w.status != 0 {
		t.Errorf("w.status = %d; expected 0", w.status)
	}
}
