package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestPusher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
	}
	if w.Pusher() != nil {
		t.Errorf("Pusher() = %v, want nil", w.Pusher())
	}
}
