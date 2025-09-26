package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionStore{}
	s.sid = "123456"
	ctx := context.Background()
	if s.SessionID(ctx) != "123456" {
		t.Errorf("SessionID() = %v, want %v", s.SessionID(ctx), "123456")
	}
}
