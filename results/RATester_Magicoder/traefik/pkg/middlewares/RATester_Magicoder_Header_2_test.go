package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rw := &ResponseModifier{req: req, rw: rr}

	header := rw.Header()

	if header == nil {
		t.Error("Expected non-nil header map; got nil")
	}
}
