package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestHijack_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	w := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
	}
	w.Hijack()
	if w.size != 0 {
		t.Errorf("w.size = %d; want 0", w.size)
	}
}
