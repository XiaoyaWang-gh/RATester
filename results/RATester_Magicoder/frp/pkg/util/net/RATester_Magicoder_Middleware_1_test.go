package net

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	authMid := &HTTPAuthMiddleware{
		user:   "testUser",
		passwd: "testPasswd",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	wrappedHandler := authMid.Middleware(handler)

	req, _ := http.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, but got %d", http.StatusUnauthorized, recorder.Code)
	}

	req.SetBasicAuth("testUser", "testPasswd")
	recorder = httptest.NewRecorder()

	wrappedHandler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}
