package session

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestSessionRelease_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	st := &CookieSessionStore{
		sid:    "test_sid",
		values: map[interface{}]interface{}{"key": "value"},
	}
	w := httptest.NewRecorder()
	st.SessionRelease(ctx, w)

	cookies := w.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(cookies))
	}
	cookie := cookies[0]
	if cookie.Name != "test_cookie_name" {
		t.Errorf("expected cookie name to be 'test_cookie_name', got '%s'", cookie.Name)
	}
	if cookie.Value != "encoded_cookie_value" {
		t.Errorf("expected cookie value to be 'encoded_cookie_value', got '%s'", cookie.Value)
	}
	if cookie.Path != "/" {
		t.Errorf("expected cookie path to be '/', got '%s'", cookie.Path)
	}
	if !cookie.HttpOnly {
		t.Error("expected cookie to be http only")
	}
	if cookie.Secure {
		t.Error("expected cookie to not be secure")
	}
	if cookie.MaxAge != 3600 {
		t.Errorf("expected cookie max age to be 3600, got %d", cookie.MaxAge)
	}
}
