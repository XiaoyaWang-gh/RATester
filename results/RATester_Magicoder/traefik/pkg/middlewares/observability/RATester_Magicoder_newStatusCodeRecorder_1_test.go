package observability

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewStatusCodeRecorder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	status := http.StatusOK
	recorder := newStatusCodeRecorder(rw, status)

	if recorder.status != status {
		t.Errorf("Expected status %d, got %d", status, recorder.status)
	}

	if recorder.ResponseWriter != rw {
		t.Errorf("Expected ResponseWriter to be %v, got %v", rw, recorder.ResponseWriter)
	}
}
