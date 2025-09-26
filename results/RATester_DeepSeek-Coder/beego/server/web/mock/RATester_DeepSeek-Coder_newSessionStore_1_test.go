package mock

import (
	"fmt"
	"testing"
)

func TestNewSessionStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := newSessionStore()

	if store.sid == "" {
		t.Errorf("Expected a non-empty session id, got: %s", store.sid)
	}

	if len(store.values) != 0 {
		t.Errorf("Expected an empty map, got: %v", store.values)
	}
}
