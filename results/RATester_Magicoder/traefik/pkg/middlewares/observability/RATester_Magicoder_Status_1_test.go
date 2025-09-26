package observability

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	recorder := &statusCodeRecorder{status: http.StatusOK}
	if recorder.Status() != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Status())
	}
}
