package memcache

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestSessionRelease_13(t *testing.T) {
	t.Run("TestSessionRelease", func(t *testing.T) {

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

		rs.values["key"] = "value"

		rs.SessionRelease(ctx, w)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
		}

		if len(rs.values) != 0 {
			t.Errorf("Expected session values to be empty, got %v", rs.values)
		}
	})
}
