package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	redirect := Redirect{
		Code:     http.StatusFound,
		Request:  req,
		Location: "/test/location",
	}

	err = redirect.Render(rr)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("Expected status code %v, got %v", http.StatusFound, status)
	}

	expectedLocation := "/test/location"
	if location := rr.Header().Get("Location"); location != expectedLocation {
		t.Errorf("Expected location header %q, got %q", expectedLocation, location)
	}
}
