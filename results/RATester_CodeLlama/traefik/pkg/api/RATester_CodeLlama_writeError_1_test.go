package api

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	writeError(rw, "test", 200)
	if rw.Code != 200 {
		t.Errorf("rw.Code = %d, want 200", rw.Code)
	}
	if rw.Body.String() != `{"Message":"test"}` {
		t.Errorf("rw.Body = %s, want `{\"Message\":\"test\"}`", rw.Body.String())
	}
}
