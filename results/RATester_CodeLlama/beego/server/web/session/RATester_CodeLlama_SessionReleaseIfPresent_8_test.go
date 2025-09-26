package session

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestSessionReleaseIfPresent_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &CookieSessionStore{}
	ctx := context.Background()
	w := httptest.NewRecorder()
	st.SessionReleaseIfPresent(ctx, w)
}
