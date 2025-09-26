package memcache

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestSessionReleaseIfPresent_13(t *testing.T) {
	t.Run("test release session", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		rs := &SessionStore{}
		ctx := context.Background()
		w := &httptest.ResponseRecorder{}
		rs.SessionReleaseIfPresent(ctx, w)
	})
}
