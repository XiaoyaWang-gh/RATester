package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoMethod_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()

	engine.NoMethod(func(c *Context) {
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	})

	server := httptest.NewServer(engine)
	defer server.Close()

	client := server.Client()

	resp, err := client.Post(server.URL, "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "Method Not Allowed" {
		t.Errorf("Expected body 'Method Not Allowed', got '%s'", string(body))
	}
}
