package context

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func Testreset_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r := &Response{
		ResponseWriter: rw,
		Started:        true,
		Status:         200,
	}

	r.reset(rw)

	if r.ResponseWriter != rw {
		t.Errorf("Expected ResponseWriter to be %v, but got %v", rw, r.ResponseWriter)
	}

	if r.Started != false {
		t.Errorf("Expected Started to be false, but got %v", r.Started)
	}

	if r.Status != 0 {
		t.Errorf("Expected Status to be 0, but got %v", r.Status)
	}
}
