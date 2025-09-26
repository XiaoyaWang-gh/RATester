package session

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSessionRelease_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &CookieSessionStore{
		sid: "123456",
		values: map[interface{}]interface{}{
			"name": "astaxie",
		},
	}
	w := httptest.NewRecorder()
	st.SessionRelease(context.Background(), w)
	cookie := w.Header().Get("Set-Cookie")
	if !strings.Contains(cookie, "name=astaxie") {
		t.Error("cookie value error")
	}
}
