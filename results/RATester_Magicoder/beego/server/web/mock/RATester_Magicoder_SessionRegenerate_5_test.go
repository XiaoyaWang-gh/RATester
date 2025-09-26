package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	oldSid := "oldSid"
	sid := "sid"
	sessionProvider := &SessionProvider{
		Store: &SessionStore{
			sid:    oldSid,
			values: make(map[interface{}]interface{}),
		},
	}

	newStore, err := sessionProvider.SessionRegenerate(ctx, oldSid, sid)
	if err != nil {
		t.Fatalf("SessionRegenerate failed: %v", err)
	}

	if newStore.SessionID(ctx) != sid {
		t.Errorf("Expected new session ID to be %s, but got %s", sid, newStore.SessionID(ctx))
	}

	if len(newStore.(*SessionStore).values) != 0 {
		t.Errorf("Expected new session store to be empty, but got %v", newStore.(*SessionStore).values)
	}
}
