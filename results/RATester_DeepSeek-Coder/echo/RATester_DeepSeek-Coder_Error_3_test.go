package echo

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestError_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := errors.New("test error")
	c.Error(err)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
	}

	expectedBody := fmt.Sprintf("{\"message\":\"%s\"}", err.Error())
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, rec.Body.String())
	}
}
