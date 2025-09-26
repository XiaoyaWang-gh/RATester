package echo

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestUnwrap_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r := &Response{Writer: rw}

	got := r.Unwrap()
	if got != rw {
		t.Errorf("Response.Unwrap() = %v, want %v", got, rw)
	}
}
