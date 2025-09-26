package session

import (
	"context"
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSessionRelease_8(t *testing.T) {
	t.Run("TestSessionRelease", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		st := &CookieSessionStore{
			sid:    "test_sid",
			values: make(map[interface{}]interface{}),
		}
		st.values["test_key"] = "test_value"

		w := httptest.NewRecorder()
		st.SessionRelease(context.Background(), w)

		cookies := w.Result().Cookies()
		if len(cookies) != 1 {
			t.Errorf("Expected 1 cookie, got %d", len(cookies))
		}

		cookie := cookies[0]
		if cookie.Name != "test_sid" {
			t.Errorf("Expected cookie name to be 'test_sid', got '%s'", cookie.Name)
		}

		decodedCookie, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			t.Errorf("Failed to unescape cookie value: %v", err)
		}

		if decodedCookie != "test_value" {
			t.Errorf("Expected decoded cookie value to be 'test_value', got '%s'", decodedCookie)
		}
	})
}
