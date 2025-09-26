package context

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r := &Response{
		ResponseWriter: rw,
		Started:        false,
	}

	data := []byte("Hello, world!")
	n, err := r.Write(data)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != len(data) {
		t.Errorf("Expected to write %d bytes, but wrote %d", len(data), n)
	}
	if r.Started != true {
		t.Error("Expected Started to be true after writing")
	}
	if rw.Body.String() != string(data) {
		t.Errorf("Expected body to be %q, but got %q", data, rw.Body.String())
	}
}
