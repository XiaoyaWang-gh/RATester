package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionStore{}
	s.sid = "test"
	ctx := context.Background()
	if s.SessionID(ctx) != "test" {
		t.Error("SessionID() failed")
	}
}
