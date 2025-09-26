package buffering

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	oxybuffer "github.com/vulcand/oxy/v2/buffer"
)

func TestServeHTTP_39(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	b := &buffer{
		name:   "test",
		buffer: &oxybuffer.Buffer{},
	}
	b.ServeHTTP(rw, req)
	if rw.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rw.Code)
	}
}
