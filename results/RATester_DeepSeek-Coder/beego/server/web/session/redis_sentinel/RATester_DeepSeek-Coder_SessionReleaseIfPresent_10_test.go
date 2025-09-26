package redis_sentinel

import (
	"context"
	"fmt"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSessionReleaseIfPresent_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := httptest.NewRecorder()

	store := &SessionStore{
		p:           redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
		sid:         "test",
		lock:        sync.RWMutex{},
		values:      make(map[interface{}]interface{}),
		maxlifetime: 3600,
	}

	store.SessionReleaseIfPresent(ctx, w)

	// Check if the session was released
	_, err := store.p.Get(ctx, "test").Result()
	if err != redis.Nil {
		t.Errorf("Expected session to be released, but it was not")
	}
}
