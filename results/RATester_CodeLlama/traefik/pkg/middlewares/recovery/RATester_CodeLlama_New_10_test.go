package recovery

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNew_10(t *testing.T) {
	t.Run("should create a new recovery middleware", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		ctx := context.Background()
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

		// when
		middleware, err := New(ctx, next)

		// then
		assert.NoError(t, err)
		assert.NotNil(t, middleware)
	})
}
