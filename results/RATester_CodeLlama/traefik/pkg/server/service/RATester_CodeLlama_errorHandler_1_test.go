package service

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorHandler_1(t *testing.T) {
	t.Run("Test errorHandler", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		err := errors.New("test error")
		errorHandler(w, req, err)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("w.Code = %d, want %d", w.Code, http.StatusInternalServerError)
		}
		if w.Body.String() != http.StatusText(http.StatusInternalServerError) {
			t.Errorf("w.Body = %s, want %s", w.Body.String(), http.StatusText(http.StatusInternalServerError))
		}
	})
}
