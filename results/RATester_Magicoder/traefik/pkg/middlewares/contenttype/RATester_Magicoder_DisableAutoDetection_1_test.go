package contenttype

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDisableAutoDetection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test logic here
	})

	testHandler := DisableAutoDetection(handler)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	testHandler.ServeHTTP(rr, req)

	// Test assertions here
}
