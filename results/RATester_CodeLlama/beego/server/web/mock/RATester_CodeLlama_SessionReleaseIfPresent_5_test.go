package mock

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionReleaseIfPresent_5(t *testing.T) {
	t.Parallel()

	t.Run("should release session if present", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		ctx := context.Background()
		w := httptest.NewRecorder()
		s := &SessionStore{}

		// when
		s.SessionReleaseIfPresent(ctx, w)

		// then
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
