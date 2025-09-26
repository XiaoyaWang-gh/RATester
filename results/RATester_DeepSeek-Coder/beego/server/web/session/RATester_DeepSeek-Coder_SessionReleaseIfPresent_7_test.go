package session

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestSessionReleaseIfPresent_7(t *testing.T) {
	t.Run("TestSessionReleaseIfPresent", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := &FileSessionStore{
			sid:    "test_sid",
			lock:   sync.RWMutex{},
			values: make(map[interface{}]interface{}),
		}

		ctx := context.Background()
		w := httptest.NewRecorder()

		fs.SessionReleaseIfPresent(ctx, w)

		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Result().StatusCode)
		}
	})
}
