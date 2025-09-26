package observability

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteHeader_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	recorder := &statusCodeRecorder{status: http.StatusOK}
	status := http.StatusInternalServerError
	recorder.WriteHeader(status)
	if recorder.status != status {
		t.Errorf("Expected status code %d, but got %d", status, recorder.status)
	}
}
