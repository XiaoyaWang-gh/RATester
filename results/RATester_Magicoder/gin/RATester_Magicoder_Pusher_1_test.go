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

	rw := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
	}

	pusher := rw.Pusher()
	if pusher != nil {
		t.Errorf("Expected nil, got %v", pusher)
	}

	rw.ResponseWriter = httptest.NewRecorder()
	pusher = rw.Pusher()
	if pusher == nil {
		t.Errorf("Expected non-nil, got nil")
	}
}
