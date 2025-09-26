package session

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSessionReleaseIfPresent_8(t *testing.T) {
	t.Run("TestSessionReleaseIfPresent", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		w := httptest.NewRecorder()
		store := &CookieSessionStore{
			sid:    "test_sid",
			values: make(map[interface{}]interface{}),
		}

		store.values["test_key"] = "test_value"

		store.SessionReleaseIfPresent(ctx, w)

		cookies := w.Result().Cookies()
		require.Equal(t, 1, len(cookies))
		require.Equal(t, "test_sid", cookies[0].Name)
		require.Equal(t, "test_sid", cookies[0].Value)
		require.True(t, cookies[0].Expires.IsZero())
		require.True(t, cookies[0].MaxAge < 0)
		require.True(t, cookies[0].Secure)
		require.True(t, cookies[0].HttpOnly)
		require.Equal(t, "Path=/; SameSite=Lax", cookies[0].Raw)
	})
}
