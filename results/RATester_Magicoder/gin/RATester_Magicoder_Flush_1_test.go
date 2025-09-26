package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestFlush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
	}

	w.Flush()

	if w.size != 0 {
		t.Errorf("Expected size to be 0 after flush, but got %d", w.size)
	}

	if w.status != 0 {
		t.Errorf("Expected status to be 0 after flush, but got %d", w.status)
	}
}
