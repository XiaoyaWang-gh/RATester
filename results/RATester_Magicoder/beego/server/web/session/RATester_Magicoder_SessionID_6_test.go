package session

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestSessionID_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	fs := &FileSessionStore{
		sid:    "test_sid",
		lock:   sync.RWMutex{},
		values: make(map[interface{}]interface{}),
	}

	expected := "test_sid"
	actual := fs.SessionID(ctx)

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
