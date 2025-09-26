package context

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r := &Response{
		ResponseWriter: rw,
		Started:        false,
		Status:         0,
	}

	r.WriteHeader(http.StatusOK)

	if r.Status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, r.Status)
	}

	if r.Started != true {
		t.Errorf("Expected Started to be true, but it was false")
	}

	if rw.Code != http.StatusOK {
		t.Errorf("Expected response writer status code %d, but got %d", http.StatusOK, rw.Code)
	}
}
