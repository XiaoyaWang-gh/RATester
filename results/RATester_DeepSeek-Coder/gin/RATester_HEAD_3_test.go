package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHEAD_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()

	router.HEAD("/path", func(c *Context) {
		c.String(200, "OK")
	})

	req, _ := http.NewRequest("HEAD", "/path", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("Expected status code 200, got %d", resp.Code)
	}

	if resp.Body.String() != "" {
		t.Errorf("Expected empty body, got %s", resp.Body.String())
	}
}
