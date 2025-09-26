package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &BasicAuth{
		Realm: "testRealm",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	auth.RequireAuth(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	expectedHeader := fmt.Sprintf("Basic realm=\"%s\"", auth.Realm)
	if resp.Header.Get("WWW-Authenticate") != expectedHeader {
		t.Errorf("Expected WWW-Authenticate header %q, got %q", expectedHeader, resp.Header.Get("WWW-Authenticate"))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "401 Unauthorized\n" {
		t.Errorf("Expected body %q, got %q", "401 Unauthorized\n", string(body))
	}
}
