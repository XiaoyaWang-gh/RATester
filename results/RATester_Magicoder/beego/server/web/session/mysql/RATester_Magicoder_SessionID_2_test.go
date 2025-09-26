package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	st := &SessionStore{
		sid: "test_session_id",
	}

	sessionID := st.SessionID(ctx)

	if sessionID != "test_session_id" {
		t.Errorf("Expected session ID to be 'test_session_id', but got '%s'", sessionID)
	}
}
