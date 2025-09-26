package render

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := httptest.NewRecorder()
	r := JSON{Data: "test"}

	err := r.Render(w)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type 'application/json', got %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	expectedBody := `"test"`
	if string(body) != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, string(body))
	}
}
