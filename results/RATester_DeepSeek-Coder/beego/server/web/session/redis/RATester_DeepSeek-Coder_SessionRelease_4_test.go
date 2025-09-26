package redis

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSessionRelease_4(t *testing.T) {
	t.Run("TestSessionRelease", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		w := httptest.NewRecorder()
		store := &SessionStore{
			p:           redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
			sid:         "test_session",
			lock:        sync.RWMutex{},
			values:      make(map[interface{}]interface{}),
			maxlifetime: 3600,
		}

		store.Set(ctx, "test_key", "test_value")
		store.SessionRelease(ctx, w)

		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Result().StatusCode)
		}
	})
}
