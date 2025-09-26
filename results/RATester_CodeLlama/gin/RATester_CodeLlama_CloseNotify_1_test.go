package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestCloseNotify_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	w := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
	}
	if w.CloseNotify() != nil {
		t.Errorf("CloseNotify() = %v, want nil", w.CloseNotify())
	}
}
