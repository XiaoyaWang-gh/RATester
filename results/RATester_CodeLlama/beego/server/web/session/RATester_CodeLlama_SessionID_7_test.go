package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &MemSessionStore{}
	st.sid = "test"
	ctx := context.Background()
	if st.SessionID(ctx) != "test" {
		t.Errorf("SessionID() = %v, want %v", st.SessionID(ctx), "test")
	}
}
