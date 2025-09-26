package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStaticFS_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()
	group := router.Group("/group")

	fs := http.Dir("./static")
	group.StaticFS("/static", fs)

	tests := []struct {
		url      string
		expected string
	}{
		{"/group/static/file1.txt", "file1 content"},
		{"/group/static/file2.txt", "file2 content"},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		if w.Body.String() != test.expected {
			t.Errorf("Expected body %s, got %s", test.expected, w.Body.String())
		}
	}
}
