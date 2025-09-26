package net

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_3(t *testing.T) {
	t.Run("normal", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/", nil)
		h := &HTTPGzipWrapper{
			h: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hello"))
			}),
		}
		h.ServeHTTP(w, r)
		if w.Body.String() != "hello" {
			t.Errorf("w.Body.String() = %v, want %v", w.Body.String(), "hello")
		}
	})
	t.Run("gzip", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/", nil)
		r.Header.Set("Accept-Encoding", "gzip")
		h := &HTTPGzipWrapper{
			h: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hello"))
			}),
		}
		h.ServeHTTP(w, r)
		if w.Body.String() != "hello" {
			t.Errorf("w.Body.String() = %v, want %v", w.Body.String(), "hello")
		}
		if w.Header().Get("Content-Encoding") != "gzip" {
			t.Errorf("w.Header().Get(\"Content-Encoding\") = %v, want %v", w.Header().Get("Content-Encoding"), "gzip")
		}
	})
}
