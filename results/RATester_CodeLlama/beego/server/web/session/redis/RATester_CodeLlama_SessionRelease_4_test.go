package redis

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestSessionRelease_4(t *testing.T) {
	t.Parallel()

	t.Run("release session", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		ctx := context.Background()
		w := httptest.NewRecorder()

		rs := &SessionStore{}
		rs.SessionRelease(ctx, w)
	})
}
