package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &FileSessionStore{}
	fs.sid = "test"
	ctx := context.Background()
	if fs.SessionID(ctx) != "test" {
		t.Error("SessionID() failed")
	}
}
