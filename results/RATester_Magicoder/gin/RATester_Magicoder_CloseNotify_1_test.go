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

	notifyChan := w.CloseNotify()

	// Test that the channel is not nil
	if notifyChan == nil {
		t.Error("Expected a non-nil channel")
	}

	// Test that the channel is closed
	select {
	case <-notifyChan:
		// Channel was closed
	default:
		t.Error("Expected the channel to be closed")
	}
}
