package session

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestSessionReleaseIfPresent_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := httptest.NewRecorder()

	store := &MemSessionStore{
		sid:          "test",
		timeAccessed: time.Now(),
		value: map[interface{}]interface{}{
			"key": "value",
		},
		lock: sync.RWMutex{},
	}

	store.SessionReleaseIfPresent(ctx, w)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	if w.Body.String() != "" {
		t.Errorf("Expected empty body, got %s", w.Body.String())
	}
}
