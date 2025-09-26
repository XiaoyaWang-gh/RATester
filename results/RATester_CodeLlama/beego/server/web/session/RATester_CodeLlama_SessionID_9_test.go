package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &CookieSessionStore{
		sid: "test",
	}
	ctx := context.Background()
	if st.SessionID(ctx) != "test" {
		t.Errorf("SessionID() = %v, want %v", st.SessionID(ctx), "test")
	}
}
