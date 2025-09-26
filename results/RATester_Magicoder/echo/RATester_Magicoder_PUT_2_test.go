package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPUT_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()

	testPath := "/test"
	testHandler := func(c Context) error {
		return nil
	}

	e.PUT(testPath, testHandler)

	req := httptest.NewRequest(http.MethodPut, testPath, nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}
