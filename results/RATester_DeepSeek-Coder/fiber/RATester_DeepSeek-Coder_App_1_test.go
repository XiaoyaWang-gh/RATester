package fiber

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Errorf("Test failed with error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}
}
