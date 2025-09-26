package fiber

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestXHR_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Requested-With", "xmlhttprequest")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
