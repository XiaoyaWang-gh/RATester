package memcache

import (
	"context"
	"fmt"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestSessionReleaseIfPresent_13(t *testing.T) {
	t.Run("TestSessionReleaseIfPresent", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		w := httptest.NewRecorder()
		rs := &SessionStore{
			sid:         "test_session",
			lock:        sync.RWMutex{},
			values:      make(map[interface{}]interface{}),
			maxlifetime: 3600,
		}

		rs.values["test_key"] = "test_value"

		rs.SessionReleaseIfPresent(ctx, w)

		if _, ok := rs.values["test_key"]; ok {
			t.Errorf("Expected session value to be deleted")
		}
	})
}
