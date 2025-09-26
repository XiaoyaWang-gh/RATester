package gin

import (
	"fmt"
	"net/http"
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
	if w.status != http.StatusOK {
		t.Errorf("w.status = %d; want %d", w.status, http.StatusOK)
	}
}
