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
		t.Errorf("SessionRegenerate() error = %v", err)
		return
	}

	if newStore.SessionID(ctx) != sid {
		t.Errorf("SessionRegenerate() new session id = %v, want %v", newStore.SessionID(ctx), sid)
	}

	if sessionProvider.Store.SessionID(ctx) != oldSid {
		t.Errorf("SessionRegenerate() old session id = %v, want %v", sessionProvider.Store.SessionID(ctx), oldSid)
	}
}
